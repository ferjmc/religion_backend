package entities

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	Type   string
	UserID uint
}
