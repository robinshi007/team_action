package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"team_action/core/user"
	"team_action/core/user/helper"

	"team_action/apps/note"
)

// init database -
func (ds *DServer) initDB() error {
	var db *gorm.DB
	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}
	// open logmode if env is dev
	db.LogMode(true)

	// postgresql only
	//db.Exec("SET search_path TO team_action_dev")
	if !db.HasTable("users") {
		db.AutoMigrate(&user.User{})
		db.AutoMigrate(
			&note.Category{},
			&note.Note{},
		)

		Password := helper.HashAndSalt([]byte("admin"))
		adminUser := user.User{
			UserName: "admin",
			Password: Password,
		}
		db.Create(&adminUser)
	}
	//db.Model(&note.Note{}).AddForeignKey("category_id", "nt_categories(id)", "CASCADE", "CASCADE")
	//sqlite3 only
	//db.Exec("PRAGMA foreign_keys = ON")

	return nil
}
