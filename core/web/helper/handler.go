package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"team_action/core/user"
	udto "team_action/core/user/dto"
)

// GetUUID -
func GetUUID(ctx *gin.Context, idString string) (uuid.UUID, error) {
	uid, err := uuid.FromString(ctx.Param(idString))
	if err != nil {
		return uuid.UUID{}, err
	}
	return uid, nil
}

// GetCurrentUser -
func GetCurrentUser(ctx *gin.Context) (*user.User, error) {
	userInterface, ok := ctx.Get(udto.IdentityKey)
	if !ok {
		return nil, errors.New("Cannot get current User")
	}
	currentUser := userInterface.(*user.User)
	return currentUser, nil
}
