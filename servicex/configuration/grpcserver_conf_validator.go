package configuration

import "errors"

//Error Validator
var (
	ErrConfGRPCServerPortNotInRange = errors.New("grpc server port configuration should between 15000 - 17000")
)

func validConfigGRPCServerPort(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	if config.GRPCServer.Port < 15000 || config.GRPCServer.Port > 17000 {
		return ErrConfGRPCServerPortNotInRange
	}
	return nil
}
