package server

import "errors"

type Validatable interface {
	Validate() error
}

var (
	ErrNotValidatable = errors.New("type is not validatable")
	ErrAPIRegistryRequire = errors.New("api registry is require")
	ErrGRPCRegistryRequire = errors.New("grpc registry is require")
	ErrGRPCRegistryServiceFuncsRequire = errors.New("grpc registry service funcs is require")


)

type Validator struct{}

func (v *Validator) Validate(i interface{}) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return ErrNotValidatable
}
