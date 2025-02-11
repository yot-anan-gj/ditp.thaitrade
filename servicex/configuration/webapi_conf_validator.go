package configuration

import (
	"errors"
)

//Error Validator
var (
	ErrConfWebApiPortNotInRange = errors.New("web api port configuration should between 5000 - 9000")

)

func validConfigWebApiPort(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	if config.WebAPI.Port < 5000 || config.WebAPI.Port > 9000 {
		return ErrConfWebApiPortNotInRange
	}
	return nil
}


