package service

import (
	"team_action/apps/note"
)

// NoteService -
type noteService struct {
	repo note.NoteRepo
}

// NewNoteService --
func NewNoteService(repo note.NoteRepo) note.NoteService {
	return &noteService{
		repo: repo,
	}
}

// Delete -
func (svc *noteService) Delete(id string) error {
	return svc.repo.Delete(id)
}

func (svc *noteService) GetAll() ([]*note.Note, error) {
	return svc.repo.GetAll()
}

func (svc *noteService) GetByID(id string) (*note.Note, error) {
	return svc.repo.GetByID(id)
}

func (svc *noteService) Store(u *note.Note) (string, error) {
	return svc.repo.Store(u)
}

func (svc *noteService) Update(u *note.Note) error {
	return svc.repo.Update(u)
}

func (svc *noteService) Search(word string) ([]*note.Note, error) {
	return svc.repo.Search(word)
}
