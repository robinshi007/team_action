package sql

import (
	"errors"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"team_action/core/config"
)

// NewDb -
func NewDb(c *config.Config) (*gorm.DB, error) {
	if c.Database.DriverName == "postgres" {
		return newPostgres(c)
	} else if c.Database.DriverName == "sqlite3" {
		return newSqlite3(c)
	}
	return nil, errors.New("Not supported db")
}
func newSqlite3(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", c.Database.URLAddress)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newPostgres(c *config.Config) (*gorm.DB, error) {
	psqlInfo := c.Database.URLAddress
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
