package webserver_handlers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/common_bindings"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_constant"
)

const WebAppDBCheckerName = "WebApp Database"

func WebAppDBHealthChecker(c echo.Context) HealthItemChecker {
	return func() *common_bindings.HealthItem {
		//ping to database
		healthItem := &common_bindings.HealthItem{
			ItemName: WebAppDBCheckerName,
		}

		//get context
		if dbConnections, ok := c.Get(server_constant.DBContextKey).(database.Connections); ok {
			errMsgs := make([]string, 0)
			successMsgs := make([]string, 0)
			for contextName, dbConn := range dbConnections {
				if dbConn == nil {
					errMsgs = append(errMsgs, fmt.Sprintf("context %s: not found connection", contextName))
				} else {
					err := dbConn.Ping()
					if err != nil {
						errMsgs = append(errMsgs, fmt.Sprintf("context %s: %s", contextName, err))
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
			healthItem.Status = UnHealthyStatus
			healthItem.Message = "not found WebApp Database"
			return healthItem
		}

	}
}
