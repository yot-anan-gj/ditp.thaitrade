package server

type ServerOption func(*Server) error

func APIRegistryOpt(apiRegistries []*APIRegistry) ServerOption {
	return func(server *Server) error {
		if len(apiRegistries) <= 0 {
			return ErrAPIRegistryRequire
		}
		server.apiRegistries = apiRegistries
		return nil
	}
}


func GRPCRegistryOpt(grpcRegistry *GRPCRegistry) ServerOption {
	return func(server *Server) error {
		if grpcRegistry == nil {
			return ErrGRPCRegistryRequire
		}

		if len(grpcRegistry.RegisterFuncs) <= 0 {
			return ErrGRPCRegistryServiceFuncsRequire
		}
		server.grpcRegistry = grpcRegistry
		return nil
	}
}
