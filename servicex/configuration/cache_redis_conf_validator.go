package configuration

import (
	"errors"
	"fmt"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrorConfRedisCacheNameReq = errors.New("redis cache name is require")

	ErrorConfRedisCacheNameDup = func(name string) error {
		return fmt.Errorf("error redis cache name %s is duplicate", name)
	}
)


func validConfigRedisCache(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	cacheNameCount := make(map[string]int)
	for _, redisCache := range config.RedisCaches{
		if stringutil.IsEmptyString(redisCache.CacheName){
			return ErrorConfRedisCacheNameReq
		}
		cacheNameCount[redisCache.CacheName]++
		if cacheNameCount[redisCache.CacheName] > 1{
			return ErrConfWebApiSessionRedisSessionNameDup(redisCache.CacheName)
		}
	}
	return nil
}

