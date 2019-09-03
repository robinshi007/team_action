package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCtrl -
type healthCtrl struct{}

// NewHealthCtrl todo
func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

// Ping -
func (h *healthCtrl) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
