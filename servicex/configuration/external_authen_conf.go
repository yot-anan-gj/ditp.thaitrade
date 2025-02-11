package configuration

import "fmt"

const (
	FACEBOOK = "facebook"
	GOOGLE   = "google"
	LINE     = "line"
	APPLE    = "apple"
)

type ExternalAuthenticationConfig struct {
	Provider        string
	AppleTeamID     string
	AppleKeyID      string
	AppleCertP8File string
	ClientID        string
	ClientSecret    string
	RedirectURL     string
	Scopes          []string
}

func (extAuth *ExternalAuthenticationConfig) String() string {
	return fmt.Sprintf(
		"Provider: %s, "+
			"AppleTeamID: %s, "+
			"AppleKeyID: %s, "+
			"AppleCertP8File: %s, "+
			"ClientID: %s, "+
			"ClientSecret: %s, "+
			"RedirectURL: %s",
		extAuth.Provider,
		extAuth.AppleTeamID,
		extAuth.AppleKeyID,
		extAuth.AppleCertP8File,
		extAuth.ClientID,
		extAuth.ClientSecret,
		extAuth.RedirectURL)
}
