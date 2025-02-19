package webserver

import "errors"

type Validatable interface {
	Validate() error
}

var (
	ErrNotValidatable = errors.New("type is not validatable")
	ErrSiteRegistryRequire = errors.New("site registry is require")
)

type Validator struct{}

func (v *Validator) Validate(i interface{}) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return ErrNotValidatable
}
