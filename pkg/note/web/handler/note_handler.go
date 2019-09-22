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
	"team_action/pkg/web"
)

type noteCtrl struct {
	log logger.LogInfoFormat
	svc note.INoteService
}

// NewNoteCtrl -
func NewNoteCtrl(log logger.LogInfoFormat, svc note.INoteService) *noteCtrl {
	return &noteCtrl{log, svc}
}

func (n *noteCtrl) GetAll(ctx *gin.Context) {
	notes, err := n.svc.GetAll()
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: notes,
	})
}

func (n *noteCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}

	note, err := n.svc.GetByID(id)
	if err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, &web.SuccessResponse{
		Data: note,
	})
}

func (n *noteCtrl) Store(ctx *gin.Context) {
	//	cid, err := uuid.FromString(ctx.PostForm("category_id"))
	//	if err != nil {
	//		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
	//		return
	//	}
	var note dto.NewNote
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		return
	}
	id, err := n.svc.Store(&ne.Note{
		CategoryID: note.CategoryID,
		Title:      note.Title,
		Body:       note.Body,
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

func (n *noteCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}

	var note dto.EditNote
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.Error(cerrors.NewParamError([]string{err.Error()}))
		//ghandler.HandleBadRequestRepsonse(err, ctx)
		return
	}
	if err := n.svc.Update(&ne.Note{
		Entity:     base.Entity{ID: uid},
		Title:      note.Title,
		Body:       note.Body,
		CategoryID: note.CategoryID,
	}); err != nil {
		ctx.Error(cerrors.NewCustomError("1103", []string{err.Error()}))
		//ghandler.HandleErrorRepsonse(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func (n *noteCtrl) Delete(ctx *gin.Context) {
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
