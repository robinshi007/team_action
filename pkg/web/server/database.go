package server

import (
	"github.com/jinzhu/gorm"

	"team_action/pkg/note"
	"team_action/pkg/user"
)

// init database -
func (ds *DServer) initDB() error {
	var db *gorm.DB
	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}
	// open logmode if env is dev
	db.LogMode(true)
	//db.Exec("SET search_path TO team_action_dev")
	db.AutoMigrate(&user.User{})

	db.AutoMigrate(
		&note.Note{},
		&note.Category{},
	)

	return nil
}
