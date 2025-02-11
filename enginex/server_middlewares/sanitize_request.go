package server_middlewares

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"html"
	"io/ioutil"
)

func SanitizeHTML() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			body := req.Body

			buf := new(bytes.Buffer)
			buf.ReadFrom(body)
			body.Close()

			p := bluemonday.UGCPolicy()

			p.AllowAttrs("style").OnElements("img","i","a","span","div","p","li","b","strong","hr","h1","h2","h3","h4","h5","label","small")

			// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
			newBody := p.SanitizeBytes(buf.Bytes())

			str := html.UnescapeString(string(newBody))

			req.Body = ioutil.NopCloser(bytes.NewReader([]byte(str)))

			c.SetRequest(req)

			return next(c)

		}
	}
}
