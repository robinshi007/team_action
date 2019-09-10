package note

// Repo - contract
type Repo interface {
	Delete(id string) error
	GetAll() ([]*Note, error)
	GetByID(id string) (*Note, error)
	Store(u *Note) (string, error)
	Update(u *Note) error
}
