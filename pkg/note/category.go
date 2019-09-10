package note

import "team_action/pkg/base"

// Category -
type Category struct {
	base.Entity
	Name        string `json:"name" gorm:"type:varchar(255);column:name;index:category_name;UNIQUE,NOT NULL"`
	Description string `json:"description" gorm:"type:text;column:description;"`
	Notes       []Note `json:"notes" gorm:"foreighkey:CategoryID"`
}

// TableName -
func (n Category) TableName() string {
	return "note_categories"
}
