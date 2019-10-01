package di

import (
	"go.uber.org/dig"

	"team_action/core/config"
	"team_action/core/logger"
	"team_action/core/sql"
	"team_action/core/user"
	user_repo "team_action/core/user/repo"

	note_repo "team_action/apps/note/repo"
	note_service "team_action/apps/note/service"
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

		// user
		container.Provide(user_repo.NewUserRepo)
		container.Provide(user.NewUserService)

		// note app
		container.Provide(note_repo.NewNoteRepo)
		container.Provide(note_service.NewNoteService)
		container.Provide(note_repo.NewCategoryRepo)
		container.Provide(note_service.NewCategoryService)

		inited = true
	}
	return container
}

// Invoke todo
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
