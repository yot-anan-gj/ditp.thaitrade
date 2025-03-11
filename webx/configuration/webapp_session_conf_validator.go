package configuration

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrorConfWebAppSessionRedisContextReq = errors.New("Redis Session Name is require")

	ErrConfWebAppSessionRedisSessionNameDup = func(sessionName string) error {
		return fmt.Errorf("error web redis session name %s is duplicate", sessionName,)
	}
)


func validConfigWebAppRedisSession(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	sessionNameCount := make(map[string]int)
	for _,redisSession := range config.WebApp.SessionStore.RedisStores{
		if stringutil.IsEmptyString(redisSession.SessionName){
			return ErrorConfWebAppSessionRedisContextReq
		}
		sessionNameCount[redisSession.SessionName]++
		if sessionNameCount[redisSession.SessionName] > 1{
			return ErrConfWebAppSessionRedisSessionNameDup(redisSession.SessionName)
		}
	}
	return nil
}
