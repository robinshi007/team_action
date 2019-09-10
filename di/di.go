package di

import (
	"team_action/pkg/config"
	"team_action/pkg/logger"
	note_repo "team_action/pkg/note/repo"
	note_service "team_action/pkg/note/service"
	"team_action/pkg/sql"
	"team_action/pkg/user"
	user_repo "team_action/pkg/user/repo"

	//	"team_action/pkg/login"

	"go.uber.org/dig"
)

// container - di container object
var container = dig.New()

// inited - check container is inited or not
var inited = false

// BuildContainer todo
func BuildContainer() *dig.Container {
	// config
	if inited == false {
		container.Provide(config.NewConfig)

		// DB
		container.Provide(sql.NewDb)

		// logger
		container.Provide(logger.NewLogger)

		// login
		// container.Provide(orm.NewLoginRepo)
		// container.Provide(login.NewLoginService)

		// user
		container.Provide(user_repo.NewUserRepo)
		container.Provide(user.NewUserService)
		// note
		container.Provide(note_repo.NewNoteRepo)
		container.Provide(note_service.NewNoteService)

		inited = true
	}
	return container
}

// Invoke todo
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
