package webapi_handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/common_bindings"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/dateutil"
)

const (
	HealthyStatus   = "healthy"
	UnHealthyStatus = "un-healthy"
)

type HealthItemChecker func() *common_bindings.HealthItem

func HealthCheck(c echo.Context) error {

	itemCheckers := []HealthItemChecker{
		WebAPIDBHealthChecker(c), WebAPICacheHealthChecker(c),
	}

	response := &common_bindings.HealthResponse{
		Status: HealthyStatus,
	}

	healthzItems := make([]*common_bindings.HealthItem, 0)

	//generate all item healthz
	for _, healthzFn := range itemCheckers {
		healthzItem := healthzFn()
		if healthzItem.Status == UnHealthyStatus && response.Status == HealthyStatus {
			//main response status
			response.Status = UnHealthyStatus
		}
		healthzItems = append(healthzItems, healthzItem)
	}

	response.Items = healthzItems
	t := time.Now()
	response.EpochTime = dateutil.DateTime2Epoch(&t)
	response.StatusTime = dateutil.DateTime2DefaultString(&t)
	if response.Status == UnHealthyStatus {
		response.StatusMessage = UnHealthyStatus
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.StatusMessage = HealthyStatus
		return c.JSON(http.StatusOK, response)
	}

}
