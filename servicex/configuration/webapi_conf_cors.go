package configuration

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var (
	defaultAllowOrigins = []string{"*"}

	defaultAllowMethods = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodPatch,
		http.MethodPost,
		http.MethodDelete,
	}

	defaultAllowHeader = []string{
		echo.HeaderOrigin,
		echo.HeaderContentType,
		echo.HeaderAccept,
		echo.HeaderAcceptEncoding,
		echo.HeaderAccessControlAllowHeaders,
		echo.HeaderAccessControlAllowMethods,
		echo.HeaderAccessControlAllowOrigin,
		echo.HeaderXRequestedWith,
		echo.HeaderAuthorization,
		echo.HeaderXCSRFToken,
	}

	defaultExposeHeaders = []string{}

	defaultMaxAge = 0
)

type WebAPICORsConfig struct {
	AllowOrigins     [] string
	AllowMethods     [] string
	AllowHeaders     [] string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
}

func (wcc WebAPICORsConfig) String() string {
	return fmt.Sprintf("AllowOrigins: %v, AllowMethods: %v, AllowHeaders: %v, AllowCredentials: %t, ExposeHeaders: %s, MaxAge: %d",
		wcc.AllowOrigins, wcc.AllowMethods,
		wcc.AllowHeaders, wcc.AllowCredentials,
		wcc.ExposeHeaders, wcc.MaxAge)

}
