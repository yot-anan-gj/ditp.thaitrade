package configuration

import (
	"errors"
	"fmt"

	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/fileutil"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var DBProviders = map[string]bool{
	//POSTGRES_AWS:        true,
	//POSTGRES_GCP:        true,
	POSTGRES_ON_PREMISE: true,
}

var (
	ErrConfWebAppDBInvalidDBProvider = func(contextName string, provider string) error {
		return fmt.Errorf("error web database configuration invalid provider %s at context name %s", provider, contextName)
	}
	ErrConfWebAppDBContextRequire = errors.New("error web database configuration context name is require")

	ErrConfWebAppDBProviderRequire = func(contextName string) error {
		return fmt.Errorf("error web database configuration provider at context name %s is require", contextName)
	}

	ErrConfWebAppDBURLRequire = func(contextName string) error {
		return fmt.Errorf("error web database configuration url at context name %s is require", contextName)
	}

	ErrConfWebAppDBUserRequire = func(contextName string) error {
		return fmt.Errorf("error web database configuration user at context name %s is require", contextName)
	}

	ErrConfWebAppDBPasswordRequire = func(contextName string) error {
		return fmt.Errorf("error web database configuration password at context name %s is require", contextName)
	}

	ErrConfWebAppDBDatabaseNameRequire = func(contextName string) error {
		return fmt.Errorf("error web database configuration password at database name %s is require", contextName)
	}

	ErrConfWebAppStaticsNotfoundLocalFolder = func(folder string) error {
		return fmt.Errorf("error web statics configuration source folder %s is not exist", folder)
	}

	ErrConfWebAppStaticsCheckLocalFolder = func(folder string, err error) error {
		return fmt.Errorf("error web statics configuration source folder %s %s", folder, err)
	}

	ErrConfWebAppStaticsLocalFolderDup = func(folder string) error {
		return fmt.Errorf("error web statics configuration source folder %s duplicate", folder)
	}

	ErrConfWebAppStaticsTargetDup = func(target string) error {
		return fmt.Errorf("error web statics configuration target path  %s duplicate", target)
	}

	ErrConfWebAppStaticsTargetRequire = errors.New("error web statics configuration target path is not blank")
)

func validConfigWebAppDB(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}
	for _, database := range config.WebApp.Databases {
		if stringutil.IsEmptyString(database.ContextName) {
			return ErrConfWebAppDBContextRequire
		}
		if stringutil.IsEmptyString(database.Provider) {
			return ErrConfWebAppDBProviderRequire(database.ContextName)
		}
		if !DBProviders[database.Provider] {
			return ErrConfWebAppDBInvalidDBProvider(database.ContextName, database.Provider)
		}
		if stringutil.IsEmptyString(database.URL) {
			return ErrConfWebAppDBURLRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.User) {
			return ErrConfWebAppDBUserRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.Password) {
			return ErrConfWebAppDBPasswordRequire(database.ContextName)
		}
		if stringutil.IsEmptyString(database.DatabaseName) {
			return ErrConfWebAppDBDatabaseNameRequire(database.ContextName)
		}

	}
	return nil
}

func validConfigWebAppStatics(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	//validate source folder
	targetCount := make(map[string]int)
	sourceCount := make(map[string]int)
	for _, staticMap := range config.WebApp.Statics {
		for target, source := range staticMap {
			targetCount[target]++
			sourceCount[source]++
			isExist, err := fileutil.IsDirExist(source)
			if err != nil {
				return ErrConfWebAppStaticsCheckLocalFolder(source, err)
			}
			if !isExist {
				return ErrConfWebAppStaticsNotfoundLocalFolder(source)
			}
			if stringutil.IsEmptyString(target) {
				return ErrConfWebAppStaticsTargetRequire
			}

		}
	}

	//valid duplicate
	for target, count := range targetCount {
		if count > 1 {
			return ErrConfWebAppStaticsTargetDup(target)
		}
	}

	for source, count := range sourceCount {
		if count > 1 {
			return ErrConfWebAppStaticsLocalFolderDup(source)
		}
	}

	return nil
}
