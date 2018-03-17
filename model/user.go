package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:30" json:"username"`
	Password string `gorm:"size:30" json:"password"`
	Nickname string `gorm:"size:50" json:"nickname"`
	Avatar string `gorm:"size:255" json:"avatar"`
}