package configuration

import (
	"errors"
	"fmt"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
	"strings"
)

var AuthenProviders = map[string]bool{
	FACEBOOK: true,
	GOOGLE: true,
	LINE: true,
	APPLE: true,
}

var(
	ErrConfExtAuthenProviderRequire = errors.New("error external authentication provider is require")
	ErrConfExtAuthenClientIDRequire = errors.New("error external authentication client id is require")
	ErrConfExtAuthenClientSecretRequire = errors.New("error external authentication client secret is require")
	ErrConfExtAuthenRedirectURLRequire = errors.New("error external authentication redirect URL is require")
	ErrorConfExtAuthenAppleTeamIDRequire = errors.New("error external authentication provider apple AppleTeamID is require")
	ErrorConfExtAuthenAppleKeyIDRequire = errors.New("error external authentication provider apple AppleKeyID is require")
	ErrorConfExtAuthenAppleCertP8FileRequire = errors.New("error external authentication provider apple AppleCertP8File is require")

	ErrConfExtAuthenInvalidAuthProvider = func(provider string) error {
		return fmt.Errorf("error external authentication invalid provider %s ", provider)
	}
	ErrConfExtAuthenDuplicateProvider = func(provider string) error {
		return fmt.Errorf("error external authentication provider %s is duplicate", provider)
	}
)


func validConfigExternalAuthen(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	providerCount := make(map[string]int)
	for _, extAuth := range config.ExternalAuthentications {
		if stringutil.IsEmptyString(extAuth.Provider) {
			return ErrConfExtAuthenProviderRequire
		}
		if !AuthenProviders[extAuth.Provider] {
			return ErrConfExtAuthenInvalidAuthProvider(extAuth.Provider)
		}

		providerCount[strings.ToLower(extAuth.Provider)]++
		if providerCount[extAuth.Provider] > 1{
			return ErrConfExtAuthenDuplicateProvider(extAuth.Provider)
		}

		if stringutil.IsEmptyString(extAuth.ClientID) {
			return ErrConfExtAuthenClientIDRequire
		}
		if stringutil.IsEmptyString(extAuth.RedirectURL) {
			return ErrConfExtAuthenRedirectURLRequire
		}

		if extAuth.Provider == APPLE{
			if stringutil.IsEmptyString(extAuth.AppleTeamID){
				return ErrorConfExtAuthenAppleTeamIDRequire
			}
			if stringutil.IsEmptyString(extAuth.AppleKeyID){
				return ErrorConfExtAuthenAppleKeyIDRequire
			}
			if stringutil.IsEmptyString(extAuth.AppleCertP8File){
				return ErrorConfExtAuthenAppleCertP8FileRequire
			}
		}else{
			if stringutil.IsEmptyString(extAuth.ClientSecret) {
				return ErrConfExtAuthenClientSecretRequire
			}
		}
	}
	return nil
}
