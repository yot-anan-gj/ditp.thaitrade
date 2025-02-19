package webserver

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_constant"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/server_middlewares"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/session"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/fileutil"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/webserver_constant"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/webserver_handlers"
	"github.com/yot-anan-gj/ditp.thaitrade/webx/webserver_middlewares"
)

var (
	ErrTemplateWebPageNameRequire = errors.New("web page template name is require")
	ErrTemplateBaseNameRequire    = errors.New("base template name is require")
	ErrTemplateNameDup            = func(name string) error { return fmt.Errorf("template name: %s is duplicate", name) }
	ErrTemplateFileRequire        = errors.New("template file is require")
	ErrTemplateFileDup            = func(templateFile string) error { return fmt.Errorf("template file: %s is duplicate", templateFile) }
	ErrTemplateFileNotExist       = func(templateFile string) error { return fmt.Errorf("template file: %s not exist", templateFile) }
	ErrTemplateFail               = func(name string, templateFile string, err error) error {
		return fmt.Errorf("error generator template %s %s %s", name, templateFile, err)
	}

	ErrSessionNameReq  = errors.New("session context appender middleware is require session name")
	ErrSessionStoreReq = errors.New("session context appender middleware is require store")
)

func (ws *WebServer) templateGenerator() (map[string]*template.Template, error) {
	if ws.website == nil {
		return nil, ErrSiteRegistryRequire
	}

	//used for validate
	nameCount := make(map[string]int)
	templateFileCount := make(map[string]int)

	//validate base template
	if len(ws.website.BaseWebPage.TemplateFiles) > 1 {
		if stringutil.IsEmptyString(ws.website.BaseWebPage.Name) {
			return nil, ErrTemplateBaseNameRequire
		}

		nameCount[ws.website.BaseWebPage.Name]++

		for _, templateFile := range ws.website.BaseWebPage.TemplateFiles {
			if stringutil.IsEmptyString(templateFile) {
				return nil, ErrTemplateFileRequire
			}

			templateFileCount[templateFile]++
			if templateFileCount[templateFile] > 1 {
				return nil, ErrTemplateFileDup(templateFile)
			}

			if isExist, _ := fileutil.IsFileExist(templateFile); !isExist {
				return nil, ErrTemplateFileNotExist(templateFile)
			}

		}
	}

	//generate base web page
	var baseTemplate *template.Template = nil
	round := -1
	for _, templateFile := range ws.website.BaseWebPage.TemplateFiles {
		round++
		rawContent, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return nil, ErrTemplateFail(ws.website.BaseWebPage.Name, templateFile, err)
		}
		if round == 0 {
			baseTemplate = template.Must(template.New(ws.website.BaseWebPage.Name).Parse(string(rawContent)))
			continue
		} else {
			if baseTemplate != nil {
				_, err = baseTemplate.Parse(string(rawContent))
				if err != nil {
					return nil, ErrTemplateFail(ws.website.BaseWebPage.Name, templateFile, err)
				}
			}

		}
	}

	//validate web page template
	for _, webPage := range ws.website.WebPages {
		if stringutil.IsEmptyString(webPage.Name) {
			return nil, ErrTemplateWebPageNameRequire
		}
		nameCount[webPage.Name]++
		if nameCount[webPage.Name] > 1 {
			return nil, ErrTemplateNameDup(webPage.Name)
		}

		for _, templateFile := range webPage.TemplateFiles {
			if stringutil.IsEmptyString(templateFile) {
				return nil, ErrTemplateFileRequire
			}
			templateFileCount[templateFile]++
			if templateFileCount[templateFile] > 1 {
				return nil, ErrTemplateFileDup(templateFile)
			}
			if isExist, _ := fileutil.IsFileExist(templateFile); !isExist {
				return nil, ErrTemplateFileNotExist(templateFile)
			}

		}
	}

	//generate base web page
	templates := make(map[string]*template.Template)
	for _, webPage := range ws.website.WebPages {

		if len(webPage.TemplateFiles) <= 0 {
			continue
		}

		var webPageTmpl *template.Template = nil
		if webPage.RequireBase && baseTemplate != nil {
			webPageTmpl = template.Must(baseTemplate.Clone())
		}
		for _, templateFile := range webPage.TemplateFiles {
			//case require base template
			rawContent, err := ioutil.ReadFile(templateFile)
			if err != nil {
				return nil, ErrTemplateFail(webPage.Name, templateFile, err)
			}

			if webPageTmpl != nil {
				_, err = webPageTmpl.Parse(string(rawContent))
				if err != nil {
					return nil, ErrTemplateFail(webPage.Name, templateFile, err)
				}

			} else {
				webPageTmpl = template.Must(template.New(webPage.Name).Parse(string(rawContent)))
			}

		}

		templates[webPage.Name] = webPageTmpl

	}

	return templates, nil
}

func skipFunc(isSkip bool) middleware.Skipper {
	return func(c echo.Context) bool {
		return isSkip
	}
}

func (ws *WebServer) handlerRegister() error {

	var baseMiddleWares = map[string]echo.MiddlewareFunc{
		server_constant.SecureMiddleware: middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "SAMEORIGIN",
			HSTSMaxAge:         3600,
		}),

		server_constant.CSRFMiddleware: server_middlewares.CSRFWithConfig(server_middlewares.CSRFConfig{
			CookieName:     ws.config.WebApp.CSRF.CookieName,
			CookieMaxAge:   ws.config.WebApp.CSRF.CookieMaxAge,
			CookieSecure:   ws.config.WebApp.CSRF.CookieSecure,
			CookieHTTPOnly: ws.config.WebApp.CSRF.CookieHTTPOnly,
			CookiePath:     ws.config.WebApp.CSRF.CookiePath,
			Skipper:        skipFunc(ws.config.WebApp.CSRF.Skip),
		}),

		server_constant.CSRFIncludeGETMethodMiddleware: server_middlewares.CSRFIncludeGETMethodWithConfig(server_middlewares.CSRFConfig{
			CookieName:     ws.config.WebApp.CSRF.CookieName,
			CookieMaxAge:   ws.config.WebApp.CSRF.CookieMaxAge,
			CookieSecure:   ws.config.WebApp.CSRF.CookieSecure,
			CookieHTTPOnly: ws.config.WebApp.CSRF.CookieHTTPOnly,
			CookiePath:     ws.config.WebApp.CSRF.CookiePath,
			TokenLookup:    server_constant.DefaultCSRFGetMethodTokenLookup,
		}),

		server_constant.CORSMiddleware: middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     ws.config.WebApp.CORs.AllowOrigins,
			AllowMethods:     ws.config.WebApp.CORs.AllowMethods,
			AllowHeaders:     ws.config.WebApp.CORs.AllowHeaders,
			AllowCredentials: ws.config.WebApp.CORs.AllowCredentials,
			ExposeHeaders:    ws.config.WebApp.CORs.ExposeHeaders,
			MaxAge:           ws.config.WebApp.CORs.MaxAge,
		}),

		server_constant.DBContextAppenderMiddleware:     server_middlewares.DBContextAppender(ws.dbConnections),
		server_constant.NoCacheMiddleware:               server_middlewares.NoCache,
		server_constant.UUIDSessionGeneratorMiddleware:  server_middlewares.UUIDGenerator(ws.sessionStores),
		webserver_constant.ViewModelGeneratorMiddleware: webserver_middlewares.ViewModelGenerator(),
	}

	defaultPageMiddleWares := []string{
		server_constant.SecureMiddleware,
		server_constant.CSRFMiddleware,
		server_constant.CORSMiddleware,
		server_constant.NoCacheMiddleware,
		server_constant.DBContextAppenderMiddleware,
		webserver_constant.ViewModelGeneratorMiddleware,
	}

	defaultAPIMiddleWares := []string{
		server_constant.SecureMiddleware,
		server_constant.CSRFMiddleware,
		server_constant.CORSMiddleware,
		server_constant.DBContextAppenderMiddleware,
	}

	if ws.website == nil {
		return ErrSiteRegistryRequire
	}

	//generate session store middleware

	sessionMiddleWares := make([]echo.MiddlewareFunc, 0)

	for sessionName, store := range ws.sessionStores {
		if stringutil.IsEmptyString(sessionName) {
			return ErrSessionNameReq
		}

		if store == nil {
			return ErrSessionStoreReq
		}

		sessionMiddleWares = append(sessionMiddleWares, session.Sessions(sessionName, store))
	}

	//register base pageAPI
	for _, pageAPI := range ws.website.BaseWebPage.PageAPIs {
		//generate base page api middleware
		apiMiddleWares := make([]echo.MiddlewareFunc, 0)
		if len(pageAPI.ServerAPIMiddleWares) <= 0 {
			for _, middleWareName := range defaultAPIMiddleWares {
				apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
			}
		} else {
			for _, middleWareName := range pageAPI.ServerAPIMiddleWares {
				apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
			}
		}
		if len(sessionMiddleWares) > 0 {
			apiMiddleWares = append(apiMiddleWares, sessionMiddleWares...)
			apiMiddleWares = append(apiMiddleWares, baseMiddleWares[server_constant.UUIDSessionGeneratorMiddleware])
		}
		apiMiddleWares = append(apiMiddleWares, pageAPI.MiddleWares...)
		ws.engine.Add(pageAPI.Method, pageAPI.URL, pageAPI.Handler, apiMiddleWares...)
	}

	//register web page handler
	for _, webPage := range ws.website.WebPages {
		pageMiddleWares := make([]echo.MiddlewareFunc, 0)
		if len(webPage.ServerPageMiddleWares) <= 0 {
			if !webPage.SkipDefaultServerAPIMiddleWares {
				for _, middleWareName := range defaultPageMiddleWares {
					pageMiddleWares = append(pageMiddleWares, baseMiddleWares[middleWareName])
				}
			}
		} else {
			for _, middleWareName := range webPage.ServerPageMiddleWares {
				pageMiddleWares = append(pageMiddleWares, baseMiddleWares[middleWareName])
			}
		}
		if len(sessionMiddleWares) > 0 {
			pageMiddleWares = append(pageMiddleWares, sessionMiddleWares...)
			pageMiddleWares = append(pageMiddleWares, baseMiddleWares[server_constant.UUIDSessionGeneratorMiddleware])
		}
		pageMiddleWares = append(pageMiddleWares, webPage.MiddleWares...)
		ws.engine.Add(webPage.Method, webPage.URL, webPage.PageHandler, pageMiddleWares...)
		if len(webPage.URLs) > 0 {
			for _, url := range webPage.URLs {
				ws.engine.Add(webPage.Method, url, webPage.PageHandler, pageMiddleWares...)
			}
		}

	}

	//register web page pageAPI
	for _, webPage := range ws.website.WebPages {
		for _, pageAPI := range webPage.PageAPIs {
			apiMiddleWares := make([]echo.MiddlewareFunc, 0)
			if len(pageAPI.ServerAPIMiddleWares) <= 0 {
				if !pageAPI.SkipDefaultServerAPIMiddleWares {
					for _, middleWareName := range defaultAPIMiddleWares {
						apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
					}
				}
			} else {
				for _, middleWareName := range pageAPI.ServerAPIMiddleWares {
					apiMiddleWares = append(apiMiddleWares, baseMiddleWares[middleWareName])
				}
			}
			if len(sessionMiddleWares) > 0 {
				apiMiddleWares = append(apiMiddleWares, sessionMiddleWares...)
				apiMiddleWares = append(apiMiddleWares, baseMiddleWares[server_constant.UUIDSessionGeneratorMiddleware])
			}
			apiMiddleWares = append(apiMiddleWares, pageAPI.MiddleWares...)
			ws.engine.Add(pageAPI.Method, pageAPI.URL, pageAPI.Handler, apiMiddleWares...)
		}
	}

	//register health check
	ws.engine.GET("/health-check",
		webserver_handlers.HealthCheck,
		baseMiddleWares[server_constant.DBContextAppenderMiddleware],
	)

	return nil
}
