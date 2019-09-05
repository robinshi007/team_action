package handler

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	u "team_action/pkg/user"
	"team_action/pkg/user/dto"
)

type helloCtrl struct{}

// NewHelloCtrl -
func NewHelloCtrl() *helloCtrl {
	return &helloCtrl{}
}

// SayHi -
func (h *helloCtrl) SayHi(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	user, _ := ctx.Get(dto.IdentityKey)
	ctx.JSON(http.StatusOK, gin.H{
		"userID":   claims[dto.IdentityKey],
		"userName": user.(*u.User).UserName,
		"text":     "Hello World.",
	})
}
