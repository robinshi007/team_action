package note

import (
	"team_action/pkg/base"

	uuid "github.com/satori/go.uuid"
)

// Note - note entity
type Note struct {
	base.Entity
	Title      string    `json:"title" gorm:"type:varchar(255);index:idx_note_title;UNIQUE,NOT NULL"`
	Body       string    `json:"body" gorm:"type:text;column:body;"`
	CategoryID uuid.UUID `json:"-" gorm:"type:uuid REFERENCES nt_categories(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Category   *Category `json:"category,omitempty" gorm:"PRELOAD:false"`
}

// TableName -
func (n Note) TableName() string {
	return "nt_notes"
}
