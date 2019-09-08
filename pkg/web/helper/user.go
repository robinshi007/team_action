package helper

import (
	"fmt"

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
		fmt.Println("usrId", user.ID)
		fmt.Println("usr pass", user.Password)
		fmt.Println("usr plain pass", pass)
		fmt.Println("compare", uhelper.ComparePassword(user.Password, []byte(pass)))
		return uhelper.ComparePassword(user.Password, []byte(pass))
	}
	return false
}

// UserIsExist -
//func UserIsExist(name string) bool {
//	var db *gorm.DB
//	d := di.BuildContainer()
//	if err := d.Invoke(func(d *gorm.DB) { db = d }); err != nil {
//		return false
//	}
//	var user u.User
//	db.Where("user_name = ?", name).First(&user)
//	if user.ID != "" {
//		return true
//	}
//	return false
//}
