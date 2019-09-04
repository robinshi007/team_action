package di

import (
	"team_action/config"
	"team_action/core/user"
	"team_action/logger"
	storage "team_action/storage/sql"
	"team_action/storage/sql/orm"

	//	"team_action/pkg/login"

	"go.uber.org/dig"
)

var container = dig.New()
var inited = false

// BuildContainer todo
func BuildContainer() *dig.Container {
	// config
	if inited == false {
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
		inited = true
	}
	return container
}

// Invoke todo
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
