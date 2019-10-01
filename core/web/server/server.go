package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"team_action/core/config"
	"team_action/core/logger"
)

// DServer -
type DServer struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

// NewServer returns new DServer
func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *DServer {
	return &DServer{
		router: e,
		cont:   c,
		logger: l,
	}
}

// Start -
func (ds *DServer) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) {
		cfg = c
	}); err != nil {
		return err
	}

	// init middeware
	ds.initMiddleware()
	// init routes
	ds.initRoutes()

	// init database
	if err := ds.initDB(); err != nil {
		return err
	}
	// init database seed

	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
