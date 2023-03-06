package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID       uint
	UserID       uint //creator
	Description  string
	TotalReplays int64
}
