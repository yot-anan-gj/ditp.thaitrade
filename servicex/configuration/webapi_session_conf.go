package configuration

import "fmt"

type WebAPISessionConfig struct {
	RedisStores [] WebApiSessionRedisConfig
}

type WebApiSessionRedisConfig struct{
	SessionName string
	RedisMaxIdle int
	RedisURL string
	RedisPassword string
	MaxAge int
	MaxLength int
	CreateConnectionTimeout int
	HttpOnly bool
	Secure bool
}

func (rdStore WebApiSessionRedisConfig) String() string {
	return fmt.Sprintf("SessionName: %s, RedisMaxIdle: %d, RedisURL: %s, RedisPassword: %s, MaxAge: %d, MaxLength: %d, CreateConnectionTimeout: %d, HttpOnly: %t, Secure: %t",
		rdStore.SessionName, rdStore.RedisMaxIdle,
		rdStore.RedisURL, rdStore.RedisPassword,
		rdStore.MaxAge, rdStore.MaxLength,
		rdStore.CreateConnectionTimeout, rdStore.HttpOnly,rdStore.Secure)
}
