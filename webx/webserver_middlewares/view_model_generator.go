package webserver_middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/configuration"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/webserver_constant"
)

func ViewModelGenerator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			vm := make(map[string]interface{})
			if csrfToken, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); ok {
				vm[webserver_constant.CsrfTokenViewModelKey] = csrfToken
			}
			if conf, err := configuration.Config(); err == nil {
				if stringutil.IsNotEmptyString(conf.WebApp.GoogleTagManager.ContainerID) {
					vm[webserver_constant.GoogleTagManagerContainerIDKey] = conf.WebApp.GoogleTagManager.ContainerID
				}
			}
			if len(vm) > 0 {
				c.Set(webserver_constant.ViewModelContextKey, vm)
			}
			return next(c)
		}
	}
}
