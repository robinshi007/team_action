package note

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Note -
type Note struct {
	ID          string    `json:"note_id"  gorm:"type:uuid;column:note_id;primary_key;"`
	Title       string    `json:"title"  gorm:"type:varchar(200);index:note_title;UNIQUE,NOT NULL"`
	Description string    `json:"description"  gorm:"type:varchar(250);column:description;"`
	Body        string    `json:"body"  gorm:"type:text;column:body;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName -
func (n Note) TableName() string {
	return "notes"
}

// BeforeCreate -
func (n *Note) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}
