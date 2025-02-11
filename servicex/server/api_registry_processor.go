package server

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
	"gitlab.com/ditp.thaitrade/enginex/server_middlewares"
	"gitlab.com/ditp.thaitrade/enginex/session"
	"gitlab.com/ditp.thaitrade/enginex/util/stringutil"
	"net/http"
)

var (
	ErrSessionNameReq  = errors.New("session context appender middleware is require session name")
	ErrSessionStoreReq = errors.New("session context appender middleware is require store")

	ErrInvalidWebAPIMethod = func(method string) error {
		return fmt.Errorf("webapi invalid http method [%s]", method)
	}

	ErrWebAPIHandlerReq = func(url string) error {
		return fmt.Errorf("webapi url [%s] handler is require", url)
	}

	ErrWebAPIUrlReq = errors.New("web API URL is require")
)

var supportHTTPMetohd = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodPatch:   true,
	http.MethodDelete:  true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

func skipFunc(isSkip bool) middleware.Skipper {
	return func(c echo.Context) bool {
		return isSkip
	}
}

func (server *Server) apiRegister() error {

	var baseMiddleWares = map[string]echo.MiddlewareFunc{
		server_constant.SecureMiddleware: middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "SAMEORIGIN",
			HSTSMaxAge:         3600,
		}),

		server_constant.CSRFMiddleware: server_middlewares.CSRFWithConfig(server_middlewares.CSRFConfig{
			CookieName:     server.config.WebAPI.CSRF.CookieName,
			CookieMaxAge:   server.config.WebAPI.CSRF.CookieMaxAge,
			CookieSecure:   server.config.WebAPI.CSRF.CookieSecure,
			CookieHTTPOnly: server.config.WebAPI.CSRF.CookieHTTPOnly,
			CookiePath:     server.config.WebAPI.CSRF.CookiePath,
			Skipper:        skipFunc(server.config.WebAPI.CSRF.Skip),
		}),

		server_constant.CORSMiddleware: middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     server.config.WebAPI.CORs.AllowOrigins,
			AllowMethods:     server.config.WebAPI.CORs.AllowMethods,
			AllowHeaders:     server.config.WebAPI.CORs.AllowHeaders,
			AllowCredentials: server.config.WebAPI.CORs.AllowCredentials,
			ExposeHeaders:    server.config.WebAPI.CORs.ExposeHeaders,
			MaxAge:           server.config.WebAPI.CORs.MaxAge,
		}),

		server_constant.DBContextAppenderMiddleware:     server_middlewares.DBContextAppender(server.dbConnections),
		server_constant.CacheStoreAppenderMiddleware:    server_middlewares.CacheStoreAppender(server.cacheStoreConnections),
		server_constant.NoCacheMiddleware:               server_middlewares.NoCache,
		server_constant.UUIDSessionGeneratorMiddleware:  server_middlewares.UUIDGenerator(server.apiSessionStores),
		server_constant.DynamoContextAppenderMiddleware: server_middlewares.DynamoContextAppender(server.dynamoDBs),
		server_constant.SqsContextAppenderMiddleware:    server_middlewares.SqsContextAppender(server.sqss),
	}

	defaultAPIMiddleWares := []string{
		server_constant.SecureMiddleware,
		server_constant.CSRFMiddleware,
		server_constant.CORSMiddleware,
		server_constant.DBContextAppenderMiddleware,
		server_constant.CacheStoreAppenderMiddleware,
		server_constant.DynamoContextAppenderMiddleware,
		server_constant.SqsContextAppenderMiddleware,
	}

	if len(server.apiRegistries) <= 0 {
		return ErrAPIRegistryRequire
	}

	//generate session store middleware

	sessionMiddleWares := make([]echo.MiddlewareFunc, 0)

	for sessionName, store := range server.apiSessionStores {
		if stringutil.IsEmptyString(sessionName) {
			return ErrSessionNameReq
		}

		if store == nil {
			return ErrSessionStoreReq
		}

		sessionMiddleWares = append(sessionMiddleWares, session.Sessions(sessionName, store))
	}

	//valid date api Registries

	for _, webAPI := range server.apiRegistries {
		if !supportHTTPMetohd[webAPI.Method] {
			return ErrInvalidWebAPIMethod(webAPI.Method)
		}
		if stringutil.IsEmptyString(webAPI.URL) {
			return ErrWebAPIUrlReq
		}
		if webAPI.Handler == nil {
			return ErrWebAPIHandlerReq(webAPI.URL)
		}

	}

	//register web page pageAPI
	for _, webAPI := range server.apiRegistries {

		apiMiddleWares := make([]echo.MiddlewareFunc, 0)
		if len(webAPI.ServerAPIMiddleWares) <= 0 {
			if !webAPI.SkipDefaultServerAPIMiddleWares {
				for _, middleWareName := range defaultAPIMiddleWares {
					apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
				}
			}
		} else {
			for _, middleWareName := range webAPI.ServerAPIMiddleWares {
				apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
			}
		}
		if len(sessionMiddleWares) > 0 {
			apiMiddleWares = append(apiMiddleWares, sessionMiddleWares...)
			apiMiddleWares = append(apiMiddleWares, baseMiddleWares[server_constant.UUIDSessionGeneratorMiddleware])
		}
		apiMiddleWares = append(apiMiddleWares, webAPI.MiddleWares...)
		server.apiServer.Add(webAPI.Method, webAPI.URL, webAPI.Handler, apiMiddleWares...)

	}

	return nil
}
