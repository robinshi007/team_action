package di

import (
	"team_action/config"
	"team_action/logger"
	"team_action/storage/sql"
	"team_action/storage/sql/orm"
	"team_action/core/user"

	//	"team_action/pkg/login"

	"go.uber.org/dig"
)

var container = dig.New()

// BuildContainer todo
func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// logger
	container.Provide(logger.NewLogger)

	// login
	// container.Provide(orm.NewLoginRepo)
	// container.Provide(login.NewLoginService)

	// user
	container.Provide(orm.NewUserRepo)
	container.Provide(user.NewUserService)
	return container
}

// Invoke todo
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
