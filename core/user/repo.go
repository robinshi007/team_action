package user

// Repository -
type Repo interface {
	Delete(id string) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Store(u *User) (string, error)
	Update(u *User) error
	UpdatePassword(u *User) error
	UpdateLastLogin(u *User) error
}
