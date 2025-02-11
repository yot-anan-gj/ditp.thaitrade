package server_middlewares

import (
	"github.com/labstack/echo"
	"gitlab.com/ditp.thaitrade/enginex/queue/aws_sqs"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
)

func SqsContextAppender(sqss aws_sqs.AwsSqss) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(sqss) > 0 {
				c.Set(server_constant.SqsContextKey, sqss)
			}

			return next(c)
		}
	}
}