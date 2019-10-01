package sql

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"team_action/core/config"
)

// NewDb -
func NewDb(c *config.Config) (*gorm.DB, error) {
	if c.DB.Use == "postgres" {
		return newPostgres(c)
	} else if c.DB.Use == "sqlite3" {
		return newSqlite3(c)
	}
	return nil, errors.New("Not supported db")
}
func newSqlite3(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", c.DB.Connection.Database)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newPostgres(c *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Connection.Host,
		c.DB.Connection.Port,
		c.DB.Connection.UserName,
		c.DB.Connection.Password,
		c.DB.Connection.Database,
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
