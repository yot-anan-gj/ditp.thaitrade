package configuration

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrorConfWebApiSessionRedisContextReq = errors.New("Redis Session Name is require")

	ErrConfWebApiSessionRedisSessionNameDup = func(sessionName string) error {
		return fmt.Errorf("error web api redis session name %s is duplicate", sessionName)
	}
)

func validConfigWebApiRedisSession(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	sessionNameCount := make(map[string]int)
	for _, redisSession := range config.WebAPI.SessionStore.RedisStores {
		if stringutil.IsEmptyString(redisSession.SessionName) {
			return ErrorConfWebApiSessionRedisContextReq
		}
		sessionNameCount[redisSession.SessionName]++
		if sessionNameCount[redisSession.SessionName] > 1 {
			return ErrConfWebApiSessionRedisSessionNameDup(redisSession.SessionName)
		}
	}
	return nil
}
