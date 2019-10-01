package base

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"team_action/pkg/user"
)

// Entity - base entity contains common columns for all tables
type Entity struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;column:id;primary_key"`
	CreatedAt   time.Time  `json:"-"`
	CreatedByID uuid.UUID  `json:"-" gorm:"type:uuid"`
	CreatedBy   *user.User `json:"-"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedByID uuid.UUID  `json:"-" gorm:"type:uuid"`
	UpdatedBy   *user.User `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedByID uuid.UUID  `json:"-" gorm:"type:uuid"`
	DeletedBy   *user.User `json:"-"`
}

// BeforeCreate will set a UUID rather then numeric ID
func (base *Entity) BeforeCreate(scope *gorm.Scope) {
	id := uuid.NewV4()
	scope.SetColumn("ID", id)
}
