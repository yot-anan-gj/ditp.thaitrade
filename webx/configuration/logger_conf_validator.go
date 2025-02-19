package configuration

import (
	"errors"
	"fmt"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
)

//Error Validator
var (
	ErrConfLogInvalidLevel = func(loglevel string) error {
		return fmt.Errorf("error log config invalid level %s", loglevel)
	}

	ErrConfLogInvalidFormat = func(logFormat string) error {
		return fmt.Errorf("error log config invalid format %s", logFormat)
	}

	ErrConfLogFormatRequire = errors.New("error log config at format is require")

	ErrConfLogLevelRequire = errors.New("error log config at level is require")
)

var LogLevels = map[string]bool{
	LogLevelDebug: true,
	LogLevelInfo:  true,
	LogLevelWarn:  true,
	LogLevelError: true,
}

var LogFormats = map[string]bool{
	LogFormatJson: true,
	LogFormatText: true,
}

func validConfigLog(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}
	if stringutil.IsEmptyString(config.Log.Level) {
		return ErrConfLogLevelRequire
	}
	if !LogLevels[config.Log.Level] {
		return ErrConfLogInvalidLevel(config.Log.Level)
	}
	if stringutil.IsEmptyString(config.Log.Format) {
		return ErrConfLogFormatRequire
	}
	if !LogFormats[config.Log.Format] {
		return ErrConfLogInvalidFormat(config.Log.Format)
	}

	return nil
}
