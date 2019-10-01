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
	"team_action/pkg/user"
	udto "team_action/pkg/user/dto"
	"team_action/pkg/web"
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
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: categories,
	})
}

// GetByID -
func (n *CategoryCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}

	category, err := n.svc.GetByID(id)
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
	currentUser, _ := ctx.Get(udto.IdentityKey)
	currentUserID := currentUser.(*user.User).ID
	//	currentUserUUID, err := uuid.FromString(currentUserID)
	//	fmt.Println("currentUserUUID", currentUserUUID)
	//	if err != nil {
	//		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
	//		return
	//	}
	var category dto.NewCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	id, err := n.svc.Store(&ne.Category{
		Entity: base.Entity{
			CreatedByID: currentUserID,
			UpdatedByID: currentUserID,
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
	id := ctx.Param("id")
	uid, err := uuid.FromString(id)
	currentUser, _ := ctx.Get(udto.IdentityKey)
	currentUserID := currentUser.(*user.User).ID
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
			UpdatedByID: currentUserID,
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
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := n.svc.Delete(id); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusNoContent)
}
