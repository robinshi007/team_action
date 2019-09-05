package helper

import (
	"github.com/jinzhu/gorm"

	"team_action/di"
	u "team_action/pkg/user"
)

// CheckAuth -
func CheckAuth(name, pass string) bool {
	var db *gorm.DB
	d := di.BuildContainer()
	if err := d.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return false
	}
	var user u.User
	db.Where("user_name = ? AND password = ?", name, pass).First(&user)
	if user.ID != "" {
		return true
	}
	return false
}
