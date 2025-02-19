package webserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database/postgres"
	log "github.com/yot-anan-gj/ditp.thaitrade/enginex/echo_logrus"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_middlewares"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/session"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/configuration"
)

var instant *WebServer

func Instant() *WebServer {
	return instant
}

type WebServer struct {
	engine        *echo.Echo
	dbConnections database.Connections
	sessionStores session.Stores
	website       *SiteRegistry
	config        *configuration.AppConfig
}

func SiteRegistryOpt(site *SiteRegistry) WebServerOption {
	return func(server *WebServer) error {
		if site == nil {
			return ErrSiteRegistryRequire
		}
		server.website = site
		return nil
	}
}

func New(opts ...WebServerOption) (*WebServer, error) {
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
	server := &WebServer{
		config: config,
	}

	server.engine = echo.New()
	server.engine.Validator = &Validator{}
	server.engine.HideBanner = true

	//common middleware

	server.engine.Logger = log.Logger()
	server.engine.Use(server_middlewares.Logger())
	server.engine.Use(middleware.Recover())
	server.engine.Use(middleware.Gzip())

	for _, setter := range opts {
		err := setter(server)
		if err != nil {
			return nil, err
		}
	}

	//setup database
	server.dbConnections = make(database.Connections)
	for _, webappDBCfg := range config.WebApp.Databases {

		maxCreateTimeout := webappDBCfg.CreateConnectionTimeout
		repeatTime := 0
	dbContextLoop:
		for {
			switch webappDBCfg.Provider {
			case configuration.POSTGRES_ON_PREMISE:
				dbConn, err := postgres.Open(webappDBCfg.URL,
					webappDBCfg.User,
					webappDBCfg.Password,
					webappDBCfg.DatabaseName)
				if dbConn != nil {
					err = dbConn.Ping()
				} else {
					err = errors.New("cannot create database connection")
				}
				if err != nil && repeatTime >= maxCreateTimeout {
					return nil, fmt.Errorf("db context name: %s %s", webappDBCfg.ContextName, err.Error())
				} else if err != nil && repeatTime < maxCreateTimeout {
					log.Logger().Warnf("db context name: %s fail %s", webappDBCfg.ContextName, err)
					time.Sleep(1 * time.Second)
					repeatTime++
					log.Logger().Warnf("db context name: %s reconnect %d time..", webappDBCfg.ContextName, repeatTime)
					continue
				}
				server.dbConnections[webappDBCfg.ContextName] = dbConn
				break dbContextLoop
				//case configuration.POSTGRES_AWS:
				//case configuration.POSTGRES_GCP:
			}
		}

	}

	//create redis session store from configuration
	server.sessionStores = make(session.Stores)
	for _, redisSession := range config.WebApp.SessionStore.RedisStores {
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

			server.sessionStores[redisSession.SessionName] = store
			break sessionLoop
		}

	}

	//route static
	for _, statics := range config.WebApp.Statics {
		for prefix, root := range statics {
			server.engine.Static(prefix, root)
		}
	}

	//validate server opt
	if server.website == nil {
		return nil, ErrSiteRegistryRequire
	}

	//template setup
	if server.website != nil {
		templates, err := server.templateGenerator()
		if err != nil {
			return nil, err
		}
		server.engine.Renderer = &TemplateRegistry{
			templates: templates,
		}

		err = server.handlerRegister()
		if err != nil {
			return nil, err
		}
	}

	instant = server
	return server, nil
}

func (ws *WebServer) Start() error {

	//K8S Zero-Downtime Rolling and gracefully shutdown

	//validate engine
	if ws.Engine() == nil {
		return errors.New("error web server engine was not created")
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
		err := ws.Engine().Start(fmt.Sprintf(":%d", config.WebApp.Port))
		if err != nil {
			if err == http.ErrServerClosed {
				log.Logger().Info("shutting down the web server ...")
			} else {
				log.Logger().Fatal(err)
			}
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

		log.Logger().Infof("health check \"/readiness\" and \"/liveness\" start on :%d", config.WebApp.HealthPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.WebApp.HealthPort), healthServerMux); err != nil {
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
	time.Sleep(time.Duration(config.WebApp.K8SZeroDownTimeThreshold) * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.WebApp.GracefulShutdownTimeout)*time.Second)
	defer cancel()

	if err := ws.engine.Shutdown(ctx); err != nil {
		log.Logger().Errorf("error shutting down server %s", err)
	} else {
		log.Logger().Info("web server gracefully shutdown")

	}

	for contextName, dbConn := range ws.dbConnections {
		if err := dbConn.Close(); err != nil {
			log.Logger().Errorf("error web server closing db context name %s %s", contextName, err)
		} else {
			log.Logger().Infof("web server database context name: %s gracefully closed", contextName)

		}
	}

	return nil
}
