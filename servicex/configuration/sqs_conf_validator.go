package configuration

import (
	"errors"

	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrConfSqsContextNameIsEmpty     = errors.New("sqs configuration check: context name must not be empty")
	ErrConfSqsRegionIsEmpty          = errors.New("sqs configuration check: region must not be empty")
	ErrConfSqsQueueNameIsEmpty       = errors.New("sqs configuration check: queue name must not be empty")
	ErrConfSqsAccessKeyIDIsEmpty     = errors.New("sqs configuration check: access key id must not be empty")
	ErrConfSqsSecretAccessKeyIsEmpty = errors.New("sqs configuration check: secret access key must not be empty")
	ErrConfSqsDuplicateContextName   = errors.New("sqs configuration check: duplicate context name")
)

func validConfigSqs(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	var mapContextName = make(map[string]int)
	for _, sqsCfg := range config.Sqss {
		if stringutil.IsEmptyString(sqsCfg.ContextName) {
			return ErrConfSqsContextNameIsEmpty
		}
		if stringutil.IsEmptyString(sqsCfg.Region) {
			return ErrConfSqsRegionIsEmpty
		}
		if len(sqsCfg.QueueName) == 0 {
			return ErrConfSqsQueueNameIsEmpty
		}
		if stringutil.IsEmptyString(sqsCfg.AccessKeyID) {
			return ErrConfSqsAccessKeyIDIsEmpty
		}
		if stringutil.IsEmptyString(sqsCfg.SecretAccessKey) {
			return ErrConfSqsSecretAccessKeyIsEmpty
		}
		mapContextName[sqsCfg.ContextName] += 1
		if mapContextName[sqsCfg.ContextName] > 1 {
			return ErrConfSqsDuplicateContextName
		}
	}
	return nil
}
