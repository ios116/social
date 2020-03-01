package cmd

import (
	"go.uber.org/dig"
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

//func Connection(master *config.DateBaseConf, slave *config.SlaveConf, tar *tarantool.Connection) *users.UserStorage {
//	connMaster, err := config.DBConnection(master)
//	if err != nil {
//		log.Fatal(err)
//	}
//	connSlave, err := config.SlaveConnection(slave)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return users.NewUserStorage(connMaster, connSlave)
//}

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
	container.Provide(users.NewUserStorage)

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
