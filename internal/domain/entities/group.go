package entities

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string
	Admin uint
	Users []*User `gorm:"many2many:user_groups;"`
}
