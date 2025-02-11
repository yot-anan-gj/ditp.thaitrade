package webapi_handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlab.com/ditp.thaitrade/enginex/database"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
	"gitlab.com/ditp.thaitrade/enginex/common_bindings"
	"strings"
)

const WebAPIDBCheckerName = "WebAPI Database"

func WebAPIDBHealthChecker(c echo.Context) HealthItemChecker {
	return func() *common_bindings.HealthItem {
		//ping to database
		healthItem := &common_bindings.HealthItem{
			ItemName: WebAPIDBCheckerName,
		}

		//get context
		if dbConnections, ok := c.Get(server_constant.DBContextKey).(database.Connections); ok {
			errMsgs := make([]string,0)
			successMsgs := make([]string,0)
			for contextName, dbConn := range dbConnections {
				if dbConn == nil {
					errMsgs = append(errMsgs,fmt.Sprintf("context %s: not found connection", contextName))
				} else {
					err := dbConn.Ping()
					if err != nil{
						errMsgs = append(errMsgs,fmt.Sprintf("context %s: %s", contextName, err))
					}else{
						successMsgs = append(successMsgs, fmt.Sprintf("context %s: ok", contextName))
					}
				}

			}
			if len(errMsgs) > 0{
				healthItem.Status = UnHealthyStatus
				msgs := make([]string,0)
				msgs = append(msgs, successMsgs...)
				msgs = append(msgs, errMsgs...)
				healthItem.Message = strings.Join(msgs, ", ")
			}else{
				healthItem.Status = HealthyStatus
				healthItem.Message = strings.Join(successMsgs, ", ")
			}

			return healthItem

		} else {
			healthItem.Status = HealthyStatus
			healthItem.Message = "not found WebApp Database"
			return healthItem
		}

	}
}
