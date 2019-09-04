package user

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User -
type User struct {
	ID          string    `json:"user_id" db:"user_id" gorm:"column:user_id;primary_key;type:uuid"`
	Email       string    `json:"email" db:"email" gorm:"varchar(200);not null,unique"`
	UserName    string    `json:"user_name" db:"user_name" gorm:"varchar(50);column:user_name;not null,unique"`
	DisplayName string    `json:"display_name" db:"display_name" gorm:"varchar(50);column:display_name;"`
	FirstName   string    `json:"first_name" db:"first_name" gorm:"varchar(50);column:first_name;"`
	LastName    string    `json:"last_name" db:"last_name" gorm:"varchar(50);column:last_name"`
	Password    string    `json:"password" db:"password" gorm:"varchar(50);not null"`
	Gender      int8      `json:"gender" db:"gender"`
	Picture     string    `json:"picture" db:"picture"`
	PhoneNumber string    `json:"phone_number" db:"phone_number" gorm:"column:phone_number"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
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
