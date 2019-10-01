package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"team_action/core/base"
	"team_action/core/cerrors"
	"team_action/core/logger"
	"team_action/core/web"
	"team_action/core/web/helper"

	"team_action/apps/note"
	ne "team_action/apps/note"
	"team_action/apps/note/dto"
)

type noteCtrl struct {
	log logger.LogInfoFormat
	svc note.NoteService
}

// NewNoteCtrl -
func NewNoteCtrl(log logger.LogInfoFormat, svc note.NoteService) *noteCtrl {
	return &noteCtrl{log, svc}
}

func (n *noteCtrl) GetAll(ctx *gin.Context) {
	notes, err := n.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: notes,
	})
}

func (n *noteCtrl) GetByID(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{"invalid id string"}))
		return
	}
	note, err := n.svc.GetByID(uid.String())
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: note,
	})
}

func (n *noteCtrl) Store(ctx *gin.Context) {
	currentUser, err := helper.GetCurrentUser(ctx)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	var note dto.NewNote
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	id, err := n.svc.Store(&ne.Note{
		Entity: base.Entity{
			CreatedByID: currentUser.ID,
			UpdatedByID: currentUser.ID,
		},
		CategoryID: note.CategoryID,
		Title:      note.Title,
		Body:       note.Body,
	})
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": id,
	})
}

func (n *noteCtrl) Update(ctx *gin.Context) {
	uid, err := helper.GetUUID(ctx, "id")
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	currentUser, err := helper.GetCurrentUser(ctx)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}

	var note dto.EditNote
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	if err := n.svc.Update(&ne.Note{
		Entity: base.Entity{
			ID:          uid,
			UpdatedByID: currentUser.ID,
		},
		Title:      note.Title,
		Body:       note.Body,
		CategoryID: note.CategoryID,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.Status(http.StatusOK)
}

func (n *noteCtrl) Delete(ctx *gin.Context) {
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

func (n *noteCtrl) Search(ctx *gin.Context) {
	word := ctx.Param("id")
	notes, err := n.svc.Search(word)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: notes,
	})
}
