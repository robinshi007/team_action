package base

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Entity - base entity contains common columns for all tables
type Entity struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;column:id;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate will set a UUID rather then numeric ID
func (base *Entity) BeforeCreate(scope *gorm.Scope) {
	id := uuid.NewV4()
	scope.SetColumn("ID", id)
}
