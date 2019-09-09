package helper

import (
	"github.com/jinzhu/gorm"

	"team_action/di"
	u "team_action/pkg/user"
	uhelper "team_action/pkg/user/helper"
)

// CheckAuth -
func CheckAuth(name, pass string) bool {
	var db *gorm.DB
	d := di.BuildContainer()
	if err := d.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return false
	}
	var user u.User
	db.Where("user_name = ?", name).First(&user)
	if user.ID != "" {
		return uhelper.ComparePassword(user.Password, []byte(pass))
	}
	return false
}
