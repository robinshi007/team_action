package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"team_action/pkg/base"
	"team_action/pkg/cerrors"
	"team_action/pkg/logger"
	"team_action/pkg/note"
	ne "team_action/pkg/note"
	"team_action/pkg/note/dto"
	"team_action/pkg/web/types"
)

// CategoryCtrl -
type CategoryCtrl struct {
	log logger.LogInfoFormat
	svc note.ICategoryService
}

// NewCategoryCtrl -
func NewCategoryCtrl(log logger.LogInfoFormat, svc note.ICategoryService) *CategoryCtrl {
	return &CategoryCtrl{log, svc}
}

// GetAll -
func (n *CategoryCtrl) GetAll(ctx *gin.Context) {
	categories, err := n.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &types.SuccessResponse{
		Data: categories,
	})
}

// GetByID -
func (n *CategoryCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}

	category, err := n.svc.GetByID(id)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &types.SuccessResponse{
		Data: category,
	})
}

// Store -
func (n *CategoryCtrl) Store(ctx *gin.Context) {
	var category dto.NewCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	id, err := n.svc.Store(&ne.Category{
		Name: category.Name,
	})
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": id,
	})
}

// Update -
func (n *CategoryCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}

	var category dto.EditCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := n.svc.Update(&ne.Category{
		Entity: base.Entity{ID: uid},
		Name:   category.Name,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

// Delete -
func (n *CategoryCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := n.svc.Delete(id); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusNoContent)
}
