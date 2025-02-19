package configuration

import "errors"

const (
	//EnvKeyServerConfigFileName : configuration file name
	EnvKeyServerConfigFileName  = "SERVER_CONFIG_NAME"
	DefaultServerConfigFileName = "server"

	EnvKeyAppSecret = "APP_SECRET_KEY"
)

type AppConfig struct {
	*Configuration
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
	}
	return singletonConfig, nil

}


var (
	ErrorInvalidConfig = errors.New("invalid configuration")
)


type ValidConfigurationFunc func(*Configuration) error


//TODO: ชื่อ context กรณีข้าม configuration  ต้องห้ามซ้ำกันยังไม่ได้ตรวจ
var validConfigValidFuncs = []ValidConfigurationFunc{
	validConfigWebAppStatics, validConfigWebAppPort, validConfigWebAppHealthzPort, validConfigWebAppDB, validConfigLog, validConfigWebAppRedisSession,
}
