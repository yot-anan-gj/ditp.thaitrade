package server_middlewares

import (
	"github.com/labstack/echo"
	"gitlab.com/ditp.thaitrade/enginex/redisstore"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
)

func CacheStoreAppender(cacheStoreConnections redisstore.CacheStoreConnections) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(cacheStoreConnections) > 0 {
				c.Set(server_constant.RedisCacheStoreKey, cacheStoreConnections)
			}

			return next(c)
		}
	}
}
