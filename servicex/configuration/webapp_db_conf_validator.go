package configuration

import (
	"errors"
	"fmt"

	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var DBProviders = map[string]bool{
	//POSTGRES_AWS:        true,
	//POSTGRES_GCP:        true,
	POSTGRES_ON_PREMISE: true,
	MYSQL_ON_PREMISE:    true,
}

var (
	ErrConfDBInvalidDBProvider = func(contextName string, provider string) error {
		return fmt.Errorf("error database configuration invalid provider %s at context name %s", provider, contextName)
	}
	ErrConfDBContextRequire = errors.New("error database configuration context name is require")

	ErrConfDBProviderRequire = func(contextName string) error {
		return fmt.Errorf("error database configuration provider at context name %s is require", contextName)
	}

	ErrConfDBURLRequire = func(contextName string) error {
		return fmt.Errorf("error database configuration url at context name %s is require", contextName)
	}

	ErrConfDBUserRequire = func(contextName string) error {
		return fmt.Errorf("error database configuration user at context name %s is require", contextName)
	}

	ErrConfDBPasswordRequire = func(contextName string) error {
		return fmt.Errorf("error database configuration password at context name %s is require", contextName)
	}

	ErrConfDBDatabaseNameRequire = func(contextName string) error {
		return fmt.Errorf("error database configuration password at database name %s is require", contextName)
	}
)

func validConfigDB(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}
	for _, database := range config.Databases {
		if stringutil.IsEmptyString(database.ContextName) {
			return ErrConfDBContextRequire
		}
		if stringutil.IsEmptyString(database.Provider) {
			return ErrConfDBProviderRequire(database.ContextName)
		}
		if !DBProviders[database.Provider] {
			return ErrConfDBInvalidDBProvider(database.ContextName, database.Provider)
		}
		if stringutil.IsEmptyString(database.URL) {
			return ErrConfDBURLRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.User) {
			return ErrConfDBUserRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.Password) {
			return ErrConfDBPasswordRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.DatabaseName) {
			return ErrConfDBDatabaseNameRequire(database.ContextName)
		}

	}
	return nil
}
