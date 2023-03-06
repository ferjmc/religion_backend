package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid        string `gorm:"index"`
	Username   string
	Name       string
	Bio        string
	Website    string
	Email      string
	ProfileUrl string
	Groups     []Group `gorm:"many2many:user_groups"`
	Followers  []Follower
	Following  []Follower
	Posts      []Post
	Comments   []Comment
}
