package entities

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string
	Admin User
	Users []*User `gorm:"many2many:user_groups;"`
}
