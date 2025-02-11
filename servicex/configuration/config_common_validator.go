package configuration

import (
	"errors"
)

//Error Validator
var (
	ErrConfHealthPortNotInRange = errors.New("web server health check port configuration should between 18000 - 18100")
)

func validConfigHealthzPort(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	if config.HealthPort < 18000 || config.HealthPort > 18100 {
		return ErrConfHealthPortNotInRange
	}
	return nil
}

