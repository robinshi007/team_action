package handler

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	u "team_action/pkg/user"
	"team_action/pkg/user/dto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
		"message":  "Hello World.",
	})
}

// Crash - test 50x handler
func (h *helloCtrl) Crash(ctx *gin.Context) {
	err := errors.New("user not found")
	if rand.Intn(10) >= 5 {
		// unexpected error
		panic("panic crash")
	} else {
		// expected error
		ctx.AbortWithStatusJSON(400, gin.H{
			"code":    err.Error(),
			"message": "expected error",
		})
		return
	}
	// never get here
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"message": "Test crash.",
	//	})
}
