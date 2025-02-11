package server_middlewares

import (
	"github.com/labstack/echo"
	"gitlab.com/ditp.thaitrade/enginex/database"
	"gitlab.com/ditp.thaitrade/enginex/database/nosql/aws_dynamodb"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
)

func DBContextAppender(dbConnections database.Connections) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(dbConnections) > 0 {
				c.Set(server_constant.DBContextKey, dbConnections)
			}

			return next(c)
		}
	}
}

func DynamoContextAppender(dynamoDBs aws_dynamodb.DynamoDBs) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(dynamoDBs) > 0 {
				c.Set(server_constant.DynamoContextKey, dynamoDBs)
			}

			return next(c)
		}
	}
}