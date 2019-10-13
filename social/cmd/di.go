package cmd

import (
	"go.uber.org/dig"
	"social/internal/config"
	"social/internal/domain/entities"
	"social/internal/domain/usecase"
	"social/internal/grpcserver"
	"social/internal/storage/users"
	"social/internal/webserver"
)

func CastToUserRepository(s *users.UserStorage) entities.UserRepository {
	return entities.UserRepository(s)
}

func CastToUseService(c *usecase.Service) usecase.UserService {
	return usecase.UserService(c)
}

func BuildContainer() *dig.Container {
	container := dig.New()
	// app config
	container.Provide(config.NewAppConf)
	// app logger
	container.Provide(config.CreateLogger)
	// DB config
	container.Provide(config.NewDateBaseConf)
	// create DB connection
	container.Provide(config.DBConnection)
	// create user storage
	container.Provide(users.NewUserStorage)
	// cast db to interface
	container.Provide(CastToUserRepository)
	// RPC config
	container.Provide(config.NewGrpcConf)
	// HTTP config
	container.Provide(config.NewHttpConf)
	// create service
	container.Provide(usecase.NewService)
	// cast use case Service to interface
	container.Provide(CastToUseService)
	// create RPC service
	container.Provide(grpcserver.NewRPCServer)
	// crate WEB server
	container.Provide(webserver.NewHttpServer)

	return container

}
