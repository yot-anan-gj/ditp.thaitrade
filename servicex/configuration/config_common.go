package configuration

import (
	"errors"
	"fmt"
	"strings"
)

const (
	//EnvKeyServerConfigFileName : configuration file name
	EnvKeyServerConfigFileName  = "SERVER_CONFIG_NAME"
	DefaultServerConfigFileName = "server"

	EnvKeyAppSecret = "APP_SECRET_KEY"
)

type AppConfig struct {
	*Configuration
	externalAuthProviders map[string]ExternalAuthenticationConfig
}

//configuration
var singletonConfig *AppConfig = nil

func Config() (*AppConfig, error) {
	if singletonConfig == nil {
		config, err := read()
		if err != nil {
			return nil, err
		}
		singletonConfig = &AppConfig{
			Configuration: config,
			externalAuthProviders: make(map[string]ExternalAuthenticationConfig),
		}
	}

	return singletonConfig, nil
}

func Reload() (*AppConfig, error) {
	config, err := read()
	if err != nil {
		return nil, err
	}
	singletonConfig = &AppConfig{
		Configuration: config,
		externalAuthProviders: make(map[string]ExternalAuthenticationConfig),
	}
	return singletonConfig, nil

}

var (
	ErrorInvalidConfig = errors.New("invalid configuration")
)

type ValidConfigurationFunc func(*Configuration) error

var validConfigValidFuncs = []ValidConfigurationFunc{
	validConfigLog, validConfigHealthzPort, validConfigExternalAuthen, validConfigGRPCServerPort,
	validConfigLog, validConfigWebApiPort, validConfigWebApiRedisSession, validConfigDB,validConfigRedisCache,
	validEmailServersConfig, validConfigDynamoDB, validConfigSqs,
}


func (appCfg *AppConfig) GetExternalAuthentication(provider string) (ExternalAuthenticationConfig, error) {
	if len(appCfg.ExternalAuthentications) > 0 && len(appCfg.externalAuthProviders) <=0 {
		//cache data
		for _, externalAuthentication := range appCfg.ExternalAuthentications {
			appCfg.externalAuthProviders[externalAuthentication.Provider] = externalAuthentication
		}
	}

	if authCfg, ok := appCfg.externalAuthProviders[strings.ToLower(provider)] ;ok{
		return authCfg, nil
	}

	return ExternalAuthenticationConfig{}, fmt.Errorf("not found provider %s",provider)
}
