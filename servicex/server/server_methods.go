package server

import (
	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/database/nosql/aws_dynamodb"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/queue/aws_sqs"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/redisstore"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/session"
	"google.golang.org/grpc"
)

func (server *Server) APIServer() *echo.Echo {
	return server.apiServer
}

func (server *Server) GRPCServer() *grpc.Server {
	return server.grpcServer
}

func (server *Server) DBConnections() database.Connections {
	return server.dbConnections

}

func (server *Server) CacheStoreConnections() redisstore.CacheStoreConnections {
	return server.cacheStoreConnections
}

func (server *Server) APISessionStores() session.Stores {
	return server.apiSessionStores
}

func (server *Server) DynamoDBs() aws_dynamodb.DynamoDBs {
	return server.dynamoDBs
}

func (server *Server) Sqss() aws_sqs.AwsSqss {
	return server.sqss
}
