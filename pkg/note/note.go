package note

import (
	"team_action/pkg/base"

	uuid "github.com/satori/go.uuid"
)

// Note - note entity
type Note struct {
	base.Entity
	Title      string    `json:"title" gorm:"type:varchar(255);index:note_title;UNIQUE,NOT NULL"`
	Body       string    `json:"body" gorm:"type:text;column:body;"`
	CategoryID uuid.UUID `json:"-" gorm:"type:uuid;column:category_id;NOT NULL"`
}

// TableName -
func (n Note) TableName() string {
	return "note_notes"
}
