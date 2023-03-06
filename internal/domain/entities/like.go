package entities

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	PostID uint
}
