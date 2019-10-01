package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"team_action/core/cerrors"
	"team_action/core/jwt"
	u "team_action/core/user"
	"team_action/core/user/dto"
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
	if rand.Intn(10) >= 5 {
		// unexpected error
		panic("panic crash")
	} else {
		// expected error
		ctx.Error(cerrors.NewCustomError("1103", []string{"Expected Error for testing"}))
		return
	}
}
