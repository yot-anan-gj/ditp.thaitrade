package configuration

import (
	"fmt"
	"os"
	"time"

	echoLog "github.com/labstack/gommon/log"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	log "github.com/yot-anan-gj/ditp.thaitrade/enginex/echo_logrus"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/cryptutil"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

// configurationReader : reading configuration file
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

	if config.WebApp.GracefulShutdownTimeout < 10 {
		config.WebApp.GracefulShutdownTimeout = 10
	} else if config.WebApp.GracefulShutdownTimeout > 60 {
		config.WebApp.GracefulShutdownTimeout = 60
	}

	if config.WebApp.K8SZeroDownTimeThreshold < 0 {
		config.WebApp.K8SZeroDownTimeThreshold = 0
	}

	config.SecretKey = secretKey

	//decrypt database user and password config
	for i := 0; i < len(config.WebApp.Databases); i++ {
		encryptUsr := config.WebApp.Databases[i].User

		//decrypt
		decryptUsr, err := cryptutil.DecryptString(encryptUsr, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt database user in context name %s, %s",
				config.WebApp.Databases[i].ContextName, err.Error())
		}

		if config.WebApp.Databases[i].CreateConnectionTimeout <= 0 {
			config.WebApp.Databases[i].CreateConnectionTimeout = 10
		}

		config.WebApp.Databases[i].User = decryptUsr

		encryptPwd := config.WebApp.Databases[i].Password

		//decrypt
		decryptPwd, err := cryptutil.DecryptString(encryptPwd, secretKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt database password in context name %s, %s",
				config.WebApp.Databases[i].ContextName,
				err.Error())
		}

		config.WebApp.Databases[i].Password = decryptPwd

	}

	//update session redis store
	for i := 0; i < len(config.WebApp.SessionStore.RedisStores); i++ {
		if config.WebApp.SessionStore.RedisStores[i].CreateConnectionTimeout <= 0 {
			config.WebApp.SessionStore.RedisStores[i].CreateConnectionTimeout = 10
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
	if len(config.WebApp.CORs.AllowOrigins) <= 0 {
		config.WebApp.CORs.AllowOrigins = defaultAllowOrigins
	}

	if len(config.WebApp.CORs.AllowMethods) <= 0 {
		config.WebApp.CORs.AllowMethods = defaultAllowMethods
	}

	if len(config.WebApp.CORs.AllowHeaders) <= 0 {
		config.WebApp.CORs.AllowHeaders = defaultAllowHeader
	}

	if len(config.WebApp.CORs.ExposeHeaders) <= 0 {
		config.WebApp.CORs.ExposeHeaders = defaultExposeHeaders
	}

	if config.WebApp.CORs.MaxAge <= 0 {
		config.WebApp.CORs.MaxAge = defaultMaxAge
	}

	//read csrf config and check
	if stringutil.IsEmptyString(config.WebApp.CSRF.CookieName) {
		config.WebApp.CSRF.CookieName = "_csrf"
	}

	if stringutil.IsEmptyString(config.WebApp.CSRF.CookiePath) {
		config.WebApp.CSRF.CookiePath = "/"
	}
	if config.WebApp.CSRF.CookieMaxAge <= 0 {
		config.WebApp.CSRF.CookieMaxAge = 86400
	}

	return config, nil

}
