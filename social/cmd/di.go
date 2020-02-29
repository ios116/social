package cmd

import (
	"go.uber.org/dig"
	"log"
	"social/internal/config"
	"social/internal/domain/entities"
	"social/internal/domain/usecase"
	"social/internal/storage/users"

	"social/internal/webserver"
)

func CastToUserRepository(s *users.UserStorage) entities.UserRepository {
	return entities.UserRepository(s)
}

func CastToUseService(c *usecase.Service) usecase.UserService {
	return usecase.UserService(c)
}

func Connection(master *config.DateBaseConf, slave *config.SlaveConf) *users.UserStorage {
	connMaster, err := config.DBConnection(master)
	if err != nil {
		log.Fatal(err)
	}
	connSlave, err := config.SlaveConnection(slave)
	if err != nil {
		log.Fatal(err)
	}
	return users.NewUserStorage(connMaster, connSlave)
}

func BuildContainer() *dig.Container {
	container := dig.New()
	// app config
	container.Provide(config.NewAppConf)
	// tarantool config
	container.Provide(config.NewTarantoolConf)
	// app logger
	container.Provide(config.CreateLogger)
	// DB config
	container.Provide(config.NewDateBaseConf)
	// slave config
	container.Provide(config.NewSlaveConf)
	// connection to date base
	container.Provide(Connection)
	// connection to tarantool
	container.Provide(config.TarantoolConnection)
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
