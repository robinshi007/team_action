package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"

	"team_action/config"
	"team_action/logger"
	"team_action/core/user"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

// NewServer returns new dserver
func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (ds *dserver) SetupDB() error {
	var db *gorm.DB
	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}
	// for postgres only
	//db.Exec("SET search_path TO team_action_dev")
	db.AutoMigrate(&user.User{})

	return nil
}

func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
