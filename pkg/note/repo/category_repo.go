package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"team_action/pkg/logger"
	"team_action/pkg/note"
)

type categoryRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

// NewCategoryRepo -
func NewCategoryRepo(db *gorm.DB, log logger.LogInfoFormat) note.ICategoryRepo {
	return &categoryRepo{db, log}
}

// CategoryIsExist -
func CategoryIsExist(db *gorm.DB, title string) bool {
	var category note.Category
	var emptyUUID = uuid.UUID{}
	db.Where("name= ?", title).First(&category)
	if category.ID != emptyUUID {
		return true
	}
	return false
}

func (u *categoryRepo) Delete(id string) error {
	u.log.Debugf("deleting the category with id : %s", id)

	if err := u.db.Delete(&note.Category{}, "id = ?", id).Error; err != nil {
		errMsg := fmt.Sprintf("[categoryRepo.Delete()] with id : %s", id)
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (u *categoryRepo) GetAll() ([]*note.Category, error) {
	u.log.Debug("get all the categories")

	categories := make([]*note.Category, 0)
	err := u.db.Order("updated_at desc").Find(&categories).Error
	if err != nil {
		return nil, errors.Wrap(err, "[categoryRepo.GetALL()]")
	}
	return categories, nil
}

func (u *categoryRepo) GetByID(id string) (*note.Category, error) {
	u.log.Debugf("get category details by id : %s", id)

	category := &note.Category{}
	err := u.db.Preload("Notes").Where("id = ?", id).First(&category).Error
	if err != nil {
		errMsg := fmt.Sprintf("[categoryRepo.GetByID()] with id : %s", id)
		return nil, errors.Wrap(err, errMsg)
	}
	return category, nil
}

func (u *categoryRepo) Store(n *note.Category) (string, error) {
	u.log.Debugf("creating the category with title: %v", n.Name)

	if CategoryIsExist(u.db, n.Name) {
		return uuid.UUID{}.String(), fmt.Errorf("[categoryRepo.Store()] category title exist: %s", n.Name)
	}
	if err := u.db.Create(&n).Error; err != nil {
		return uuid.UUID{}.String(), errors.Wrap(err, "[categoryRepo.Store()] error when creating the category")
	}
	return n.ID.String(), nil
}

func (u *categoryRepo) Update(n *note.Category) error {
	u.log.Debugf("updating the category, id : %v", n.ID)

	err := u.db.Model(&n).Updates(note.Category{Name: n.Name}).Error
	if err != nil {
		errMsg := fmt.Sprintf("[categoryRepo.Update()] error while updating the category")
		return errors.Wrap(err, errMsg)
	}
	return nil
}
