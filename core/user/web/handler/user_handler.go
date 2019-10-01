package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"team_action/core/cerrors"
	"team_action/core/logger"
	ue "team_action/core/user"
	"team_action/core/user/dto"
	"team_action/core/web"
	"team_action/core/web/helper"
)

type userCtrl struct {
	log logger.LogInfoFormat
	svc ue.Service
}

// NewUserCtrl -
func NewUserCtrl(log logger.LogInfoFormat, svc ue.Service) *userCtrl {
	return &userCtrl{log, svc}
}

func (u *userCtrl) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: users,
	})
}

func (u *userCtrl) GetByID(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	user, err := u.svc.GetByID(uid.String())
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: user,
	})
}

func (u *userCtrl) Store(ctx *gin.Context) {
	var user dto.NewUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	id, err := u.svc.Store(&ue.User{
		UserName: user.UserName,
		Password: user.Password,
	})
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": id,
	})
}

func (u *userCtrl) Update(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	var user dto.NewUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := u.svc.Update(&ue.User{
		ID:       uid,
		UserName: user.UserName,
		Password: user.Password,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusOK)
}

func (u *userCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := u.svc.Delete(id); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (u *userCtrl) UpdatePassword(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	currentUser, err := helper.GetCurrentUser(ctx)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	userCopy, err := u.svc.GetByID(uid.String())
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	// only user self and admin can change the password
	if currentUser.UserName != userCopy.UserName && currentUser.UserName != "admin" {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}

	var user dto.EditPasswordUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := u.svc.Update(&ue.User{
		ID:       uid,
		Password: user.Password,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusOK)
}
func (u *userCtrl) UpdateLastLogin(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.FromString(id)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := u.svc.UpdateLastLogin(&ue.User{
		ID:          uid,
		LastLoginAt: time.Now(),
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusOK)
}
