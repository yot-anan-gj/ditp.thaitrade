package configuration

import (
	"errors"

	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrConfDynamoContextNameIsEmpty     = errors.New("dynamoDB configuration check: context name must not be empty")
	ErrConfDynamoRegionIsEmpty          = errors.New("dynamoDB configuration check: region must not be empty")
	ErrConfDynamoTableNameIsEmpty       = errors.New("dynamoDB configuration check: table name must not be empty")
	ErrConfDynamoAccessKeyIDIsEmpty     = errors.New("dynamoDB configuration check: access key id must not be empty")
	ErrConfDynamoSecretAccessKeyIsEmpty = errors.New("dynamoDB configuration check: secret access key must not be empty")
	ErrConfDynamoDuplicateContextName   = errors.New("dynamoDB configuration check: duplicate context name")
)

func validConfigDynamoDB(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	var mapContextName = make(map[string]int)
	for _, dynamoCfg := range config.DynamoDBs {
		if stringutil.IsEmptyString(dynamoCfg.ContextName) {
			return ErrConfDynamoContextNameIsEmpty
		}
		if stringutil.IsEmptyString(dynamoCfg.Region) {
			return ErrConfDynamoRegionIsEmpty
		}
		//if stringutil.IsEmptyString(dynamoCfg.TableName) {
		//	return ErrConfDynamoTableNameIsEmpty
		//}
		if stringutil.IsEmptyString(dynamoCfg.AccessKeyID) {
			return ErrConfDynamoAccessKeyIDIsEmpty
		}
		if stringutil.IsEmptyString(dynamoCfg.SecretAccessKey) {
			return ErrConfDynamoSecretAccessKeyIsEmpty
		}
		mapContextName[dynamoCfg.ContextName] += 1
		if mapContextName[dynamoCfg.ContextName] > 1 {
			return ErrConfDynamoDuplicateContextName
		}
	}
	return nil
}
