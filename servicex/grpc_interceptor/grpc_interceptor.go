package grpc_interceptor

import (
	"gitlab.com/ditp.thaitrade/enginex/database"
	"gitlab.com/ditp.thaitrade/enginex/database/nosql/aws_dynamodb"
	"gitlab.com/ditp.thaitrade/enginex/queue/aws_sqs"
	"gitlab.com/ditp.thaitrade/enginex/redisstore"
	"gitlab.com/ditp.thaitrade/enginex/server_constant"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(dbConns database.Connections, cacheCons redisstore.CacheStoreConnections, dynamoDBs aws_dynamodb.DynamoDBs, sqss aws_sqs.AwsSqss) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = context.WithValue(ctx, server_constant.DBContextKey, dbConns)
		ctx = context.WithValue(ctx, server_constant.RedisCacheStoreKey, cacheCons)
		ctx = context.WithValue(ctx, server_constant.DynamoContextKey, dynamoDBs)
		ctx = context.WithValue(ctx, server_constant.SqsContextKey, sqss)
		return handler(ctx, req)
	}
}

func StreamServerInterceptor(dbConns database.Connections, cacheCons redisstore.CacheStoreConnections, dynamoDBs aws_dynamodb.DynamoDBs, sqss aws_sqs.AwsSqss) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := WrapServerStream(ss)
		var ctx context.Context
		if wrapped.Context() == nil {
			ctx = context.Background()
		} else {
			ctx = wrapped.Context()
		}
		ctx = context.WithValue(ctx, server_constant.DBContextKey, dbConns)
		ctx = context.WithValue(ctx, server_constant.DynamoContextKey, dynamoDBs)
		ctx = context.WithValue(ctx, server_constant.SqsContextKey, sqss)
		wrapped.WrappedContext = context.WithValue(ctx, server_constant.RedisCacheStoreKey, cacheCons)

		return handler(srv, wrapped)
	}
}
