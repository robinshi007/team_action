package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"team_action/pkg/logger"
	"team_action/pkg/user"
)

type userRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

// NewUserRepo -
func NewUserRepo(db *gorm.DB, log logger.LogInfoFormat) user.Repo {
	return &userRepo{db, log}
}

// UserIsExist -
func UserIsExist(db *gorm.DB, name string) bool {
	var user user.User
	db.Where("user_name = ?", name).First(&user)
	if user.ID != "" {
		return true
	}
	return false
}

func (u *userRepo) Delete(id string) error {
	u.log.Debugf("deleting the user with id : %s", id)

	if err := u.db.Delete(&user.User{}, "user_id = ?", id).Error; err != nil {
		errMsg := fmt.Sprintf("[userRepo.Delete()] with id : %s", id)
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (u *userRepo) GetAll() ([]*user.User, error) {
	u.log.Debug("get all the users")

	users := make([]*user.User, 0)
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "[userRepo.GetALL()]")
	}
	return users, nil
}

func (u *userRepo) GetByID(id string) (*user.User, error) {
	u.log.Debugf("get user details by id : %s", id)

	user := &user.User{}
	err := u.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		errMsg := fmt.Sprintf("[userRepo.GetByID()] with id : %s", id)
		return nil, errors.Wrap(err, errMsg)
	}
	return user, nil
}

func (u *userRepo) Store(usr *user.User) error {
	u.log.Debugf("creating the user with email : %v", usr.Email)

	if UserIsExist(u.db, usr.UserName) {
		return fmt.Errorf("[userRepo.Store()] User name exist: %s", usr.UserName)
	}
	err := u.db.Create(&usr).Error
	if err != nil {
		errMsg := fmt.Sprintf("[userRepo.Store()] error when creating the user")
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (u *userRepo) Update(usr *user.User) error {
	u.log.Debugf("updating the user, user_id : %v", usr.ID)

	err := u.db.Model(&usr).Updates(user.User{UserName: usr.UserName, FirstName: usr.FirstName, LastName: usr.LastName, Password: usr.Password, Picture: usr.Picture, PhoneNumber: usr.PhoneNumber}).Error
	if err != nil {
		errMsg := fmt.Sprintf("[userRepo.Update()] error while updating the user")
		return errors.Wrap(err, errMsg)
	}
	return nil
}
