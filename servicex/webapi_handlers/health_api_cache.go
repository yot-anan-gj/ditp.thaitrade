package webapi_handlers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/common_bindings"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/redisstore"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_constant"
)

const WebAPIDCacheCheckerName = "WebAPI Cache"

func WebAPICacheHealthChecker(c echo.Context) HealthItemChecker {
	return func() *common_bindings.HealthItem {
		//ping to cache
		healthItem := &common_bindings.HealthItem{
			ItemName: WebAPIDCacheCheckerName,
		}

		//get context
		if cacheConnections, ok := c.Get(server_constant.RedisCacheStoreKey).(redisstore.CacheStoreConnections); ok {
			errMsgs := make([]string, 0)
			successMsgs := make([]string, 0)
			for contextName, cacheConn := range cacheConnections {
				if cacheConn == nil {
					errMsgs = append(errMsgs, fmt.Sprintf("context %s: not found connection", contextName))
				} else {
					bool, err := cacheConn.Ping()
					if err != nil {
						errMsgs = append(errMsgs, fmt.Sprintf("context %s: %s", contextName, err))
					} else if err == nil && !bool {
						errMsgs = append(errMsgs, fmt.Sprintf("context %s: no response", contextName))
					} else {
						successMsgs = append(successMsgs, fmt.Sprintf("context %s: ok", contextName))
					}
				}

			}
			if len(errMsgs) > 0 {
				healthItem.Status = UnHealthyStatus
				msgs := make([]string, 0)
				msgs = append(msgs, successMsgs...)
				msgs = append(msgs, errMsgs...)
				healthItem.Message = strings.Join(msgs, ", ")
			} else {
				healthItem.Status = HealthyStatus
				healthItem.Message = strings.Join(successMsgs, ", ")
			}

			return healthItem

		} else {
			healthItem.Status = HealthyStatus
			healthItem.Message = "not found Redis cache store"
			return healthItem
		}

	}
}
