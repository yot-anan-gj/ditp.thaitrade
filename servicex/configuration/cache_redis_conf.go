package configuration

import "fmt"

type RedisCacheConfig struct{
	CacheName string
	RedisMaxIdle int
	RedisURL string
	RedisPassword string
	MaxLength int
	CreateConnectionTimeout int
}

func (cache *RedisCacheConfig) String() string {
	return fmt.Sprintf("CacheName: %s, RedisMaxIdle: %d, RedisURL: %s, RedisPassword: %s, MaxLength: %d, CreateConnectionTimeout: %d",
		cache.CacheName, cache.RedisMaxIdle,
		cache.RedisURL, cache.RedisPassword,
		cache.MaxLength, cache.CreateConnectionTimeout)
}
