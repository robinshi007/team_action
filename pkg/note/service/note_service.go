package service

import (
	"team_action/pkg/note"
)

// NoteService -
type NoteService struct {
	repo note.INoteRepo
}

// NewNoteService --
func NewNoteService(repo note.INoteRepo) note.INoteService {
	return &NoteService{
		repo: repo,
	}
}

// Delete -
func (svc *NoteService) Delete(id string) error {
	return svc.repo.Delete(id)
}

func (svc *NoteService) GetAll() ([]*note.Note, error) {
	return svc.repo.GetAll()
}

func (svc *NoteService) GetByID(id string) (*note.Note, error) {
	return svc.repo.GetByID(id)
}

func (svc *NoteService) Store(u *note.Note) (string, error) {
	return svc.repo.Store(u)
}

func (svc *NoteService) Update(u *note.Note) error {
	return svc.repo.Update(u)
}
