package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"

	"team_action/pkg/config"
	"team_action/pkg/logger"
	"team_action/pkg/user"
	"team_action/pkg/web/handler"
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

func (ds *dserver) InitMiddleware() {
	// setup global middeware
	ds.router.Use(gin.Logger())
	ds.router.Use(gin.Recovery())
	ds.router.Use(handler.InternalServerErrRecover())
}

func (ds *dserver) InitDB() error {
	var db *gorm.DB
	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}
	// open logmode if env is dev
	db.LogMode(true)
	//db.Exec("SET search_path TO team_action_dev")
	db.AutoMigrate(&user.User{})
	db.Model(&user.User{}).AddUniqueIndex("idx_user_name", "user_name")

	return nil
}

func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
