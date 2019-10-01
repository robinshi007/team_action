package server

import (
	"github.com/gin-gonic/gin"

	"team_action/core/web/handler"
)

// InitMiddleware -
func (ds *DServer) initMiddleware() {
	// setup global middeware
	ds.router.Use(gin.Logger())
	ds.router.Use(gin.Recovery())
	ds.router.Use(handler.ErrorRecover())
	ds.router.Use(handler.ErrorHandling())
}
