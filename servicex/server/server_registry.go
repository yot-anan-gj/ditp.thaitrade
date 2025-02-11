package server

import (
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

type APIRegistry struct {
	URL                             string
	Handler                         echo.HandlerFunc
	Method                          string
	SkipDefaultServerAPIMiddleWares bool
	ServerAPIMiddleWares            []string
	MiddleWares                     []echo.MiddlewareFunc
}

type GRPCRegistry struct {
	ServerOptions []grpc.ServerOption
	RegisterFuncs []GRPCServiceRegisterFunc
}

type GRPCServiceRegisterFunc func(grpcServer *grpc.Server) error
