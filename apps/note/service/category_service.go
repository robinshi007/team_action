package service

import (
	"team_action/apps/note"
)

// CategoryService -
type categoryService struct {
	repo note.CategoryRepo
}

// NewCategoryService -
func NewCategoryService(repo note.CategoryRepo) note.CategoryService {
	return &categoryService{
		repo: repo,
	}
}

// Delete -
func (svc *categoryService) Delete(id string) error {
	return svc.repo.Delete(id)
}

func (svc *categoryService) GetAll() ([]*note.Category, error) {
	return svc.repo.GetAll()
}

func (svc *categoryService) GetByID(id string) (*note.Category, error) {
	return svc.repo.GetByID(id)
}

func (svc *categoryService) Store(u *note.Category) (string, error) {
	return svc.repo.Store(u)
}

func (svc *categoryService) Update(u *note.Category) error {
	return svc.repo.Update(u)
}
