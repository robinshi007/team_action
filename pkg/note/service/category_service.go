package service

import (
	"team_action/pkg/note"
)

// CategoryService -
type CategoryService struct {
	repo note.ICategoryRepo
}

// NewCategoryService -
func NewCategoryService(repo note.ICategoryRepo) note.ICategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// Delete -
func (svc *CategoryService) Delete(id string) error {
	return svc.repo.Delete(id)
}

func (svc *CategoryService) GetAll() ([]*note.Category, error) {
	return svc.repo.GetAll()
}

func (svc *CategoryService) GetByID(id string) (*note.Category, error) {
	return svc.repo.GetByID(id)
}

func (svc *CategoryService) Store(u *note.Category) (string, error) {
	return svc.repo.Store(u)
}

func (svc *CategoryService) Update(u *note.Category) error {
	return svc.repo.Update(u)
}
