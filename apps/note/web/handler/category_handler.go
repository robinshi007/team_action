package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"team_action/core/base"
	"team_action/core/cerrors"
	"team_action/core/logger"
	"team_action/core/web"
	"team_action/core/web/helper"

	"team_action/apps/note"
	ne "team_action/apps/note"
	"team_action/apps/note/dto"
)

// CategoryCtrl -
type CategoryCtrl struct {
	log logger.LogInfoFormat
	svc note.CategoryService
}

// NewCategoryCtrl -
func NewCategoryCtrl(log logger.LogInfoFormat, svc note.CategoryService) *CategoryCtrl {
	return &CategoryCtrl{log, svc}
}

// GetAll -
func (n *CategoryCtrl) GetAll(ctx *gin.Context) {
	categories, err := n.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: categories,
	})
}

// GetByID -
func (n *CategoryCtrl) GetByID(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	category, err := n.svc.GetByID(uid.String())
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: category,
	})
}

// GetByName -
func (n *CategoryCtrl) GetByName(ctx *gin.Context) {
	name := ctx.Param("id")

	category, err := n.svc.GetByName(name)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: category,
	})
}

// Store -
func (n *CategoryCtrl) Store(ctx *gin.Context) {
	currentUser, err := helper.GetCurrentUser(ctx)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	var category dto.NewCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	id, err := n.svc.Store(&ne.Category{
		Entity: base.Entity{
			CreatedByID: currentUser.ID,
			UpdatedByID: currentUser.ID,
		},
		Name: category.Name,
	})
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": id,
	})
}

// Update -
func (n *CategoryCtrl) Update(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{"Invalid id string"}))
		return
	}

	currentUser, err := helper.GetCurrentUser(ctx)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	var category dto.EditCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := n.svc.Update(&ne.Category{
		Entity: base.Entity{
			ID:          uid,
			UpdatedByID: currentUser.ID,
		},
		Name: category.Name,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusOK)
}

// Delete -
func (n *CategoryCtrl) Delete(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := n.svc.Delete(uid.String()); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusNoContent)
}
