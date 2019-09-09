package user

import "team_action/pkg/user/helper"

// Service -
type Service interface {
	Delete(id string) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Store(u *User) (string, error)
	Update(u *User) error
}

type userService struct {
	repo Repo
}

// NewUserService -
func NewUserService(repo Repo) Service {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) Delete(id string) error {
	return svc.repo.Delete(id)
}

func (svc *userService) GetAll() ([]*User, error) {
	return svc.repo.GetAll()
}

func (svc *userService) GetByID(id string) (*User, error) {
	return svc.repo.GetByID(id)
}

func (svc *userService) Store(u *User) (string, error) {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	return svc.repo.Store(u)
}

func (svc *userService) Update(u *User) error {
	return svc.repo.Update(u)
}
