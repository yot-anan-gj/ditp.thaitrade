package configuration

import (
	"fmt"
	echoLog "github.com/labstack/gommon/log"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	log "gitlab.com/ditp.thaitrade/enginex/echo_logrus"
	"gitlab.com/ditp.thaitrade/enginex/util/cryptutil"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
	"os"
	"time"
)

//configurationReader : reading configuration file
func read() (*Configuration, error) {
	vp := viper.New()
	vp.AutomaticEnv()
	//default is from configuration/config_constant.go
	vp.SetDefault(EnvKeyServerConfigFileName, DefaultServerConfigFileName)
	configName := vp.GetString(EnvKeyServerConfigFileName)
	secretKey := vp.GetString(EnvKeyAppSecret)

	if stringutil.IsEmptyString(secretKey) {
		return nil, fmt.Errorf("error environment APP_SECRET_KEY is require")
	}

	vp.SetConfigName(configName)
	vp.AddConfigPath("conf")

	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file fail, %s", err.Error())
	}

	config := &Configuration{}
	err := vp.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("error unable to decode into struct, %s", err.Error())
	}

	//validate all configuration
	for _, validFunc := range validConfigValidFuncs {
		if err := validFunc(config); err != nil {
			return nil, err
		}
	}

	if config.GracefulShutdownAPITimeout < 10 {
		config.GracefulShutdownAPITimeout = 10
	} else if config.GracefulShutdownAPITimeout > 60 {
		config.GracefulShutdownAPITimeout = 60
	}

	if config.K8SZeroDownTimeThreshold < 0 {
		config.K8SZeroDownTimeThreshold = 0
	}

	config.SecretKey = secretKey

	//decrypt database user and password config
	for i := 0; i < len(config.Databases); i++ {
		encryptUsr := config.Databases[i].User

		//decrypt
		decryptUsr, err := cryptutil.DecryptString(encryptUsr, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt database user in context name %s, %s",
				config.Databases[i].ContextName, err.Error())
		}

		if config.Databases[i].CreateConnectionTimeout <= 0 {
			config.Databases[i].CreateConnectionTimeout = 20
		}

		config.Databases[i].User = decryptUsr

		encryptPwd := config.Databases[i].Password

		//decrypt
		decryptPwd, err := cryptutil.DecryptString(encryptPwd, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt database password in context name %s, %s",
				config.Databases[i].ContextName,
				err.Error())
		}

		config.Databases[i].Password = decryptPwd

	}

	//decrypt smtp server users & password
	for i := 0; i < len(config.EmailServers.SMTPs); i++ {
		encryptUser := config.EmailServers.SMTPs[i].User
		encryptPassword:= config.EmailServers.SMTPs[i].Password

		var decryptUsr string
		if stringutil.IsNotEmptyString(encryptUser) {
			decryptUsr, err = cryptutil.DecryptString(encryptUser, secretKey)
			if err != nil {
				return nil, fmt.Errorf("unable to decrypt smtp user in server name %s, %s",
					config.EmailServers.SMTPs[i].Name, err.Error())
			}
		}

		var decryptPassword string
		if stringutil.IsNotEmptyString(encryptPassword) {
			decryptPassword, err = cryptutil.DecryptString(encryptPassword, secretKey)
			if err != nil {
				return nil, fmt.Errorf("unable to decrypt smtp password in server name %s, %s",
					config.EmailServers.SMTPs[i].Name, err.Error())
			}
		}

		config.EmailServers.SMTPs[i].User = decryptUsr
		config.EmailServers.SMTPs[i].Password = decryptPassword

	}

	//decrypt awsses secret users & password
	for i := 0; i < len(config.EmailServers.AWSSES); i++ {
		accessKeyID := config.EmailServers.AWSSES[i].AWSAccessKeyID
		secretAccessKey:= config.EmailServers.AWSSES[i].AWSSecretAccessKey

		decryptAccessKey, err := cryptutil.DecryptString(accessKeyID, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt awsses access key in server name %s, %s",
				config.EmailServers.AWSSES[i].Name, err.Error())
		}

		decryptSecretAccessKey, err := cryptutil.DecryptString(secretAccessKey, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt awsses secret access key in server name %s, %s",
				config.EmailServers.AWSSES[i].Name, err.Error())
		}

		config.EmailServers.AWSSES[i].AWSAccessKeyID = decryptAccessKey
		config.EmailServers.AWSSES[i].AWSSecretAccessKey = decryptSecretAccessKey

	}

	//update session redis store
	for i := 0; i < len(config.WebAPI.SessionStore.RedisStores); i++ {
		if config.WebAPI.SessionStore.RedisStores[i].CreateConnectionTimeout <= 0 {
			config.WebAPI.SessionStore.RedisStores[i].CreateConnectionTimeout = 20
		}
	}

	//update session redis store
	for i := 0; i < len(config.RedisCaches); i++ {
		if config.RedisCaches[i].CreateConnectionTimeout <= 0 {
			config.RedisCaches[i].CreateConnectionTimeout = 20
		}
	}

	//setup logger from config
	//log level
	switch config.Log.Level {
	case LogLevelDebug:
		log.Logger().SetLevel(echoLog.DEBUG)
	case LogLevelInfo:
		log.Logger().SetLevel(echoLog.INFO)
	case LogLevelWarn:
		log.Logger().SetLevel(echoLog.WARN)
	case LogLevelError:
		log.Logger().SetLevel(echoLog.ERROR)
	default:
		log.Logger().SetLevel(echoLog.INFO)
	}

	switch config.Log.Format {
	case LogFormatText:
		log.Logger().SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
		log.Logger().SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		})
	case LogFormatJson:
		log.Logger().SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	default:
		log.Logger().SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
		log.Logger().SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: time.RFC3339,
		})
	}

	//read cors config and check
	if len(config.WebAPI.CORs.AllowOrigins) <= 0 {
		config.WebAPI.CORs.AllowOrigins = defaultAllowOrigins
	}

	if len(config.WebAPI.CORs.AllowMethods) <= 0 {
		config.WebAPI.CORs.AllowMethods = defaultAllowMethods
	}

	if len(config.WebAPI.CORs.AllowHeaders) <= 0 {
		config.WebAPI.CORs.AllowHeaders = defaultAllowHeader
	}

	if len(config.WebAPI.CORs.ExposeHeaders) <= 0 {
		config.WebAPI.CORs.ExposeHeaders = defaultExposeHeaders
	}

	if config.WebAPI.CORs.MaxAge <= 0 {
		config.WebAPI.CORs.MaxAge = defaultMaxAge
	}

	//read csrf config and check
	if stringutil.IsEmptyString(config.WebAPI.CSRF.CookieName) {
		config.WebAPI.CSRF.CookieName = "_csrf"
	}

	if stringutil.IsEmptyString(config.WebAPI.CSRF.CookiePath) {
		config.WebAPI.CSRF.CookiePath = "/"
	}
	if config.WebAPI.CSRF.CookieMaxAge <= 0 {
		config.WebAPI.CSRF.CookieMaxAge = 86400
	}

	return config, nil

}
