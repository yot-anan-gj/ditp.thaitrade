package configuration

import (
	"errors"
)

//Error Validator
var (
	ErrConfWebAppPortNotInRange = errors.New("web server port configuration should between 5000 - 9000")
	ErrConfWebAppHealthPortNotInRange = errors.New("web server health check port configuration should between 18000 - 18100")

)

func validConfigWebAppPort(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	if config.WebApp.Port < 5000 || config.WebApp.Port > 9000 {
		return ErrConfWebAppPortNotInRange
	}
	return nil
}


func validConfigWebAppHealthzPort(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	if config.WebApp.HealthPort < 18000 || config.WebApp.HealthPort > 18100 {
		return ErrConfWebAppHealthPortNotInRange
	}
	return nil
}

