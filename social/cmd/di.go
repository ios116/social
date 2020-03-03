package cmd

import (
	"go.uber.org/dig"
	"social/internal/config"
	"social/internal/domain/entities"
	"social/internal/domain/usecase"
	"social/internal/storage"
	"social/internal/webserver"
)

func CastToUserRepository(s *storage.Storage) entities.UserRepository {
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

	// tarantool config
	container.Provide(config.NewTarantoolConf)
	container.Provide(config.TarantoolConnection)

	// DB config
	container.Provide(config.NewDateBaseConf)
	container.Provide(config.DBConnection)

	// slave config
	container.Provide(config.NewSlaveConf)
	container.Provide(config.SlaveConnection)

	// connection to date base
	container.Provide(storage.NewStorage)

	// cast db to interface
	container.Provide(CastToUserRepository)
	// HTTP config
	container.Provide(config.NewHttpConf)
	// create service
	container.Provide(usecase.NewService)
	// cast use case Service to interface
	container.Provide(CastToUseService)
	// crate WEB server
	container.Provide(webserver.NewHttpServer)

	return container

}
