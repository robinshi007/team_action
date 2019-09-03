package handler

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"team_action/core/user"
	u "team_action/core/user"
)

type helloCtrl struct{}

// NewHelloCtrl -
func NewHelloCtrl() *helloCtrl {
	return &helloCtrl{}
}

// SayHi -
func (h *helloCtrl) SayHi(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	user, _ := ctx.Get(user.IdentityKey)
	ctx.JSON(http.StatusOK, gin.H{
		"userID":   claims[u.IdentityKey],
		"userName": user.(*u.Dto).UserName,
		"text":     "Hello World.",
	})
}
