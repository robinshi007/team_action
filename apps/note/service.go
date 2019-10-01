package note

// NoteService - contract
type NoteService interface {
	Delete(id string) error
	GetAll() ([]*Note, error)
	GetByID(id string) (*Note, error)
	Store(u *Note) (string, error)
	Update(u *Note) error
	Search(word string) ([]*Note, error)
}

// CategoryService -
type CategoryService interface {
	Delete(id string) error
	GetAll() ([]*Category, error)
	GetByID(id string) (*Category, error)
	Store(u *Category) (string, error)
	Update(u *Category) error
}
