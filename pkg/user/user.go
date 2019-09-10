package user

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User -
type User struct {
	ID          string    `json:"user_id"  gorm:"type:uuid;column:user_id;primary_key;"`
	Email       string    `json:"-"  gorm:"type:varchar(200);UNIQUE,NOT NULL"`
	UserName    string    `json:"username"  gorm:"type:varchar(50);column:user_name;index:user_name;unique_index:user_name;UNIQUE,NOT NULL"`
	DisplayName string    `json:"-"  gorm:"type:varchar(50);column:display_name;"`
	FirstName   string    `json:"-"  gorm:"type:varchar(50);column:first_name;"`
	LastName    string    `json:"-" gorm:"type:varchar(50);column:last_name"`
	Password    string    `json:"-" gorm:"type:varchar(50);NOT NULL"`
	Gender      int8      `json:"-"`
	Picture     string    `json:"-"`
	PhoneNumber string    `json:"-" gorm:"column:phone_number"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName -
func (User) TableName() string {
	return "users"
}

// BeforeCreate -
func (u *User) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}
