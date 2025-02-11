package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/ditp.thaitrade/enginex/database"
	"gitlab.com/ditp.thaitrade/enginex/database/nosql/aws_dynamodb"
	"gitlab.com/ditp.thaitrade/enginex/database/postgres"
	log "gitlab.com/ditp.thaitrade/enginex/echo_logrus"
	"gitlab.com/ditp.thaitrade/enginex/queue/aws_sqs"
	"gitlab.com/ditp.thaitrade/enginex/redisstore"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
	"gitlab.com/ditp.thaitrade/enginex/server_middlewares"
	"gitlab.com/ditp.thaitrade/enginex/session"
	"gitlab.com/ditp.thaitrade/enginex/util/rdsutil"
	"gitlab.com/ditp.thaitrade/servicex/configuration"
	"gitlab.com/ditp.thaitrade/servicex/grpc_interceptor"
	"gitlab.com/ditp.thaitrade/servicex/webapi_handlers"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

var instant *Server

func Instant() *Server {
	return instant
}

type Server struct {
	apiServer             *echo.Echo
	grpcServer            *grpc.Server
	dbConnections         database.Connections
	cacheStoreConnections redisstore.CacheStoreConnections
	dynamoDBs             aws_dynamodb.DynamoDBs
	sqss                  aws_sqs.AwsSqss
	apiSessionStores      session.Stores
	apiRegistries         []*APIRegistry
	grpcRegistry          *GRPCRegistry
	config                *configuration.AppConfig
}

func New(opts ...ServerOption) (*Server, error) {
	if instant != nil {
		err := errors.New("server is already instantiate cannot re-instantiate again")
		log.Warnf("%s", err)
		return instant, err
	}
	//force load config
	config, err := configuration.Config()
	if err != nil {
		return nil, err
	}
	server := &Server{
		config: config,
	}

	server.apiServer = echo.New()
	server.apiServer.Validator = &Validator{}
	server.apiServer.HideBanner = true

	//common middleware

	server.apiServer.Logger = log.Logger()
	server.apiServer.Use(server_middlewares.Logger())
	server.apiServer.Use(middleware.Recover())

	for _, setter := range opts {
		err := setter(server)
		if err != nil {
			return nil, err
		}
	}

	//setup database
	server.dbConnections = make(database.Connections)
	for _, db := range config.Databases {

		maxCreateTimeout := db.CreateConnectionTimeout
		repeatTime := 0
	dbContextLoop:
		for {
			switch db.Provider {
			case configuration.POSTGRES_ON_PREMISE:
				dbConn, err := postgres.Open(db.URL,
					db.User,
					db.Password,
					db.DatabaseName)
				if dbConn != nil {
					err = dbConn.Ping()
				} else {
					err = errors.New("cannot create database connection")
				}
				if err != nil && repeatTime >= maxCreateTimeout {
					return nil, fmt.Errorf("db context name: %s %s", db.ContextName, err.Error())
				} else if err != nil && repeatTime < maxCreateTimeout {
					log.Logger().Warnf("db context name: %s fail %s", db.ContextName, err)
					time.Sleep(1 * time.Second)
					repeatTime++
					log.Logger().Warnf("db context name: %s reconnect %d time..", db.ContextName, repeatTime)
					continue
				}
				server.dbConnections[db.ContextName] = dbConn
				//execute initial script
				if len(db.InitialScripts) > 0 {
					err := dbInit(db.InitialScripts, db.ContextName, dbConn)
					if err != nil {
						return nil, fmt.Errorf("db context name: %s initial script fail: %s", db.ContextName, err.Error())
					}
				}
				break dbContextLoop
				//case configuration.POSTGRES_AWS:
				//case configuration.POSTGRES_GCP:
			case configuration.MYSQL_ON_PREMISE:
				dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&autocommit=true",
					db.User,
					db.Password,
					db.URL,
					db.DatabaseName))
				if dbConn != nil {
					err = dbConn.Ping()
				} else {
					err = errors.New("cannot create database connection")
				}
				if err != nil && repeatTime >= maxCreateTimeout {
					return nil, fmt.Errorf("db context name: %s %s", db.ContextName, err.Error())
				} else if err != nil && repeatTime < maxCreateTimeout {
					log.Logger().Warnf("db context name: %s fail %s", db.ContextName, err)
					time.Sleep(1 * time.Second)
					repeatTime++
					log.Logger().Warnf("db context name: %s reconnect %d time..", db.ContextName, repeatTime)
					continue
				}
				server.dbConnections[db.ContextName] = dbConn
				//execute initial script
				if len(db.InitialScripts) > 0 {
					err := dbInit(db.InitialScripts, db.ContextName, dbConn)
					if err != nil {
						return nil, fmt.Errorf("db context name: %s initial script fail: %s", db.ContextName, err.Error())
					}
				}
				break dbContextLoop
			}
		}

	}

	//setup dynamodb
	server.dynamoDBs = make(aws_dynamodb.DynamoDBs)
	for _, dynamoDB := range config.DynamoDBs {
		config, err := dynamoDB.GetAwsConfiguration(config.SecretKey)
		if err != nil {
			return nil, fmt.Errorf("dynamoDB GetAwsConfiguration fail: %s", err.Error())
		}
		db, err := aws_dynamodb.New(config...)
		server.dynamoDBs[dynamoDB.ContextName] = db
	}

	//setup sqs
	server.sqss = make(aws_sqs.AwsSqss)
	for _, sqs := range config.Sqss {
		config, err := sqs.GetAwsConfiguration(config.SecretKey)
		if err != nil {
			return nil, fmt.Errorf("sqs GetAwsConfiguration fail: %s", err.Error())
		}
		qe, err := aws_sqs.New(config...)
		server.sqss[sqs.ContextName] = qe
	}

	//create redis session store from configuration
	server.apiSessionStores = make(session.Stores)
	for _, redisSession := range config.WebAPI.SessionStore.RedisStores {
		if redisSession.RedisMaxIdle < 10 {
			redisSession.RedisMaxIdle = 10
		}

		var redisMaxAge int
		var redisMaxLength int
		//set max age
		if redisSession.MaxAge*60 < 60 || redisSession.MaxAge*60 > 24*60*60 {
			redisMaxAge = 60
		} else {
			redisMaxAge = redisSession.MaxAge * 60
		}

		//set max length
		if redisSession.MaxLength*1024 < 4*1024 || redisSession.MaxLength*1024 > 250*1024*1024 {
			redisMaxLength = 4 * 1024
		} else {
			redisMaxLength = redisSession.MaxLength * 1024 * 1024
		}

		maxCreateTimeout := redisSession.CreateConnectionTimeout
		repeatTime := 0

	sessionLoop:
		for {

			var store session.RedisStore
			var err error

			if redisSession.HttpOnly || redisSession.Secure {
				store, err = session.NewRedisStoreWithSecret(redisSession.RedisMaxIdle,
					redisMaxAge, redisMaxLength,
					"tcp", redisSession.RedisURL,
					redisSession.RedisPassword, redisSession.HttpOnly, redisSession.Secure, []byte(config.SecretKey))

			} else {
				store, err = session.NewRedisStore(redisSession.RedisMaxIdle,
					redisMaxAge, redisMaxLength,
					"tcp", redisSession.RedisURL,
					redisSession.RedisPassword, []byte(config.SecretKey))

			}

			if err != nil && repeatTime >= maxCreateTimeout {
				return nil, fmt.Errorf("redis session name: %s %s", redisSession.SessionName, err)
			} else if err != nil && repeatTime < maxCreateTimeout {
				log.Logger().Warnf("redis session name: %s fail %s", redisSession.SessionName, err)
				time.Sleep(1 * time.Second)
				repeatTime++
				log.Logger().Warnf("redis session name: %s reconnect %d time..", redisSession.SessionName, repeatTime)
				continue
			}

			server.apiSessionStores[redisSession.SessionName] = store
			break sessionLoop
		}

	}

	//create redis cache configuration
	redisCacheStoreConns := make(redisstore.CacheStoreConnections)
	for _, cache := range config.RedisCaches {
		if cache.RedisMaxIdle < 10 {
			cache.RedisMaxIdle = 10
		}

		var redisCacheMaxLength int

		//set max length
		if cache.MaxLength*1024 < 4*1024 || cache.MaxLength*1024 > 250*1024*1024 {
			redisCacheMaxLength = 4 * 1024
		} else {
			redisCacheMaxLength = cache.MaxLength * 1024 * 1024
		}

		maxCreateTimeout := cache.CreateConnectionTimeout
		repeatTime := 0

	cacheLoop:
		for {

			cacheStore, err := redisstore.NewRedisCacheStore(
				cache.RedisMaxIdle,
				redisCacheMaxLength,
				"tcp",
				cache.RedisURL,
				cache.RedisPassword)

			if err != nil && repeatTime >= maxCreateTimeout {
				return nil, fmt.Errorf("redis cache name: %s %s", cache.CacheName, err)
			} else if err != nil && repeatTime < maxCreateTimeout {
				log.Logger().Warnf("redis cache name: %s fail %s", cache.CacheName, err)
				time.Sleep(1 * time.Second)
				repeatTime++
				log.Logger().Warnf("redis cache name: %s reconnect %d time..", cache.CacheName, repeatTime)
				continue
			}

			redisCacheStoreConns[cache.CacheName] = cacheStore
			break cacheLoop
		}

	}
	if len(redisCacheStoreConns) > 0 {
		server.cacheStoreConnections = redisCacheStoreConns
	}

	//add health-check api
	//register health check
	healthCheckRegistry := &APIRegistry{
		URL:     "/health-check",
		Handler: webapi_handlers.HealthCheck,
		Method:  http.MethodGet,
		ServerAPIMiddleWares: []string{
			server_constant.DBContextAppenderMiddleware,
			server_constant.CacheStoreAppenderMiddleware,
		},
		MiddleWares: nil,
	}

	//validate server opt
	if len(server.apiRegistries) <= 0 {
		server.apiRegistries = make([]*APIRegistry, 0)
		server.apiRegistries = append(server.apiRegistries, healthCheckRegistry)
	} else {
		server.apiRegistries = append(server.apiRegistries, healthCheckRegistry)
	}

	//template setup
	if server.apiRegistries != nil {
		err = server.apiRegister()
		if err != nil {
			return nil, err
		}
	}

	instant = server
	return server, nil
}

func (server *Server) Start() error {

	//K8S Zero-Downtime Rolling and gracefully shutdown

	//validate APIServer
	if server.apiServer == nil {
		return errors.New("error server api server was not created")
	}

	//user for health check handler
	var (
		ready   = true
		muReady sync.RWMutex
	)

	//read config
	config, err := configuration.Config()
	if err != nil {
		return err
	}

	go func() {
		err := server.apiServer.Start(fmt.Sprintf(":%d", config.WebAPI.Port))
		if err != nil {
			if err == http.ErrServerClosed {
				log.Logger().Info("shutting down the api server ...")
			} else {
				log.Logger().Fatal(err)
			}
		}
	}()

	//grpc server start
	go func() {
		if server.grpcRegistry != nil && len(server.grpcRegistry.RegisterFuncs) > 0 {
			grpcPort := config.GRPCServer.Port
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
			if err != nil {
				log.Logger().Fatal(err)
			}

			grpcServerOptions := make([]grpc.ServerOption, 0)
			//generate default interceptor
			grpcServerOptions = append(grpcServerOptions,
				grpc.UnaryInterceptor(grpc_interceptor.UnaryServerInterceptor(server.dbConnections, server.cacheStoreConnections, server.dynamoDBs, server.sqss)),
				grpc.StreamInterceptor(grpc_interceptor.StreamServerInterceptor(server.dbConnections, server.cacheStoreConnections, server.dynamoDBs, server.sqss)),
			)
			if len(server.grpcRegistry.ServerOptions) > 0 {
				grpcServerOptions = append(grpcServerOptions, server.grpcRegistry.ServerOptions...)
			}
			grpcServer := grpc.NewServer(grpcServerOptions...)
			for _, registFunc := range server.grpcRegistry.RegisterFuncs {
				if err := registFunc(grpcServer); err != nil {
					log.Logger().Fatal(err)
				}
			}
			server.grpcServer = grpcServer
			log.Logger().Infof("GRPC server start at :%d", grpcPort)
			grpcServer.Serve(lis)
		}
	}()

	// health check
	go func() {
		healthServerMux := http.NewServeMux()

		healthServerMux.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
			muReady.RLock()
			checkReady := ready
			muReady.RUnlock()
			if checkReady {
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusServiceUnavailable)
		})

		healthServerMux.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		log.Logger().Infof("health check \"/readiness\" and \"/liveness\" start on :%d", config.HealthPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HealthPort), healthServerMux); err != nil {
			log.Logger().Fatalf("can not start health check server %s", err)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, os.Interrupt)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGKILL)

	<-gracefulStop

	muReady.Lock()
	ready = false
	muReady.Unlock()
	time.Sleep(time.Duration(config.K8SZeroDownTimeThreshold) * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GracefulShutdownAPITimeout)*time.Second)
	defer cancel()

	if err := server.apiServer.Shutdown(ctx); err != nil {
		log.Logger().Errorf("error shutting down server %s", err)
	} else {
		log.Logger().Info("web server gracefully shutdown")
	}

	if server.grpcServer != nil {
		server.grpcServer.GracefulStop()
		log.Logger().Info("GRPC server gracefully shutdown")
	}

	for contextName, dbConn := range server.dbConnections {
		if err := dbConn.Close(); err != nil {
			log.Logger().Errorf("error server closing db context name %s %s", contextName, err)
		} else {
			log.Logger().Infof("server database context name: %s gracefully closed", contextName)

		}
	}

	for cacheName, cache := range server.cacheStoreConnections {
		if err := cache.Close(); err != nil {
			log.Logger().Errorf("error server closing cache name %s %s", cacheName, err)
		} else {
			log.Logger().Infof("server cache name: %s gracefully closed", cacheName)

		}
	}

	return nil
}

func dbInit(initialScripts []string, contextName string, conn *sql.DB) error {
	if conn == nil {
		return fmt.Errorf("not found database connection at contextname %s", contextName)
	}

	//load sql file
	for _, initialScript := range initialScripts {
		sqlCmd, err := rdsutil.SQLLoader(initialScript)
		if err != nil {
			return err
		}
		if _, err := conn.Exec(sqlCmd); err != nil {
			return err
		}

	}

	return nil
}
