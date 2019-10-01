package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"team_action/core/logger"

	"team_action/apps/note"
)

type noteRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

// NewNoteRepo -
func NewNoteRepo(db *gorm.DB, log logger.LogInfoFormat) note.NoteRepo {
	return &noteRepo{db, log}
}

// NoteIsExist -
func NoteIsExist(db *gorm.DB, title string) bool {
	var note note.Note
	var emptyUUID = uuid.UUID{}
	db.Where("title= ?", title).First(&note)
	if note.ID != emptyUUID {
		return true
	}
	return false
}

func (u *noteRepo) Delete(id string) error {
	u.log.Debugf("deleting the note with id : %s", id)

	if err := u.db.Delete(&note.Note{}, "id = ?", id).Error; err != nil {
		errMsg := fmt.Sprintf("[noteRepo.Delete()] with id : %s", id)
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (u *noteRepo) GetAll() ([]*note.Note, error) {
	u.log.Debug("get all the notes")

	notes := make([]*note.Note, 0)
	err := u.db.Preload("UpdatedBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, user_name")
	}).Preload("Category").Order("updated_at desc").Find(&notes).Error
	if err != nil {
		return nil, errors.Wrap(err, "[noteRepo.GetALL()]")
	}
	return notes, nil
}

func (u *noteRepo) GetByID(id string) (*note.Note, error) {
	u.log.Debugf("get note details by id : %s", id)

	note := &note.Note{}
	err := u.db.Preload("UpdatedBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, user_name")
	}).Preload("Category").Where("id = ?", id).First(&note).Error
	if err != nil {
		errMsg := fmt.Sprintf("[noteRepo.GetByID()] with id : %s", id)
		return nil, errors.Wrap(err, errMsg)
	}
	return note, nil
}

func (u *noteRepo) Store(n *note.Note) (string, error) {
	u.log.Debugf("creating the note with title: %v", n.Title)

	if NoteIsExist(u.db, n.Title) {
		return uuid.UUID{}.String(), fmt.Errorf("[noteRepo.Store()] note title exist: %s", n.Title)
	}
	if err := u.db.Create(&n).Error; err != nil {
		return uuid.UUID{}.String(), errors.Wrap(err, "[noteRepo.Store()] error when creating the note")
	}
	return n.ID.String(), nil
}

func (u *noteRepo) Update(n *note.Note) error {
	u.log.Debugf("updating the note, id : %v", n.ID)

	err := u.db.Model(&n).Updates(note.Note{Title: n.Title, Body: n.Body, CategoryID: n.CategoryID}).Error
	if err != nil {
		errMsg := fmt.Sprintf("[noteRepo.Update()] error while updating the note")
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (u *noteRepo) Search(word string) ([]*note.Note, error) {
	u.log.Debug("get all the notes")

	notes := make([]*note.Note, 0)
	err := u.db.Preload("Category").Where("title like ?", fmt.Sprintf("%%%s%%", word)).Or("body like ?", fmt.Sprintf("%%%s%%", word)).Find(&notes).Error
	if err != nil {
		return nil, errors.Wrap(err, "[noteRepo.Search()]")
	}
	return notes, nil
}
