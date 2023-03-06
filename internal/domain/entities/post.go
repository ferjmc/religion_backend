package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID       uint
	Description  string
	PostImageUrl string
	Likes        []Like
	Comments     []Comment
	Book         uint32
	VerseBegin   uint32
	VerseEnd     uint32
}
