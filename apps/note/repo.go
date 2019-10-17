package note

// NoteRepo - contract
type NoteRepo interface {
	Delete(id string) error
	GetAll() ([]*Note, error)
	GetByID(id string) (*Note, error)
	Store(u *Note) (string, error)
	Update(u *Note) error
	Search(word string) ([]*Note, error)
}

// CategoryRepo - contract
type CategoryRepo interface {
	Delete(id string) error
	GetAll() ([]*Category, error)
	GetByID(id string) (*Category, error)
	GetByName(name string) (*Category, error)
	Store(u *Category) (string, error)
	Update(u *Category) error
}
