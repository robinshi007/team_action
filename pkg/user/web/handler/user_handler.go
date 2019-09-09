package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"team_action/pkg/cerrors"
	"team_action/pkg/logger"
	"team_action/pkg/user"
	up "team_action/pkg/user"
	"team_action/pkg/user/dto"
	"team_action/pkg/web/types"
)

type userCtrl struct {
	log logger.LogInfoFormat
	svc user.Service
}

// NewUserCtrl -
func NewUserCtrl(log logger.LogInfoFormat, svc user.Service) *userCtrl {
	return &userCtrl{log, svc}
}

func (u *userCtrl) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &types.SuccessResponse{
		Data: users,
	})
}

func (u *userCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}

	user, err := u.svc.GetByID(id)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &types.SuccessResponse{
		Data: user,
	})
}

func (u *userCtrl) Store(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := u.svc.Store(&up.User{
		UserName: user.UserName,
		Password: user.Password,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.Status(http.StatusCreated)
}

func (u *userCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}

	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := u.svc.Update(&up.User{
		ID:       id,
		UserName: user.UserName,
		Password: user.Password,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func (u *userCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := u.svc.Delete(id); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusNoContent)
}
