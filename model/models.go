package model

import "github.com/jinzhu/gorm"

//Item 商品
type Item struct {
	gorm.Model
	//Title 商品名
	Title string `gorm:"size:100" json:"title"`
	//Price 价格
	Price uint `json:"price"`
	//Reserve 余量
	Reserve int `json:"reserve"`
	//Cover image
	Cover string `gorm:"size:255" json:"cover"`
}

//ItemInfo object
type ItemInfo struct {
	gorm.Model
	ItemID  uint   `json:"itemId"`
	Content string `gorm:"size:3000" json:"content"`
}

//User object
type User struct {
	gorm.Model
	//Username for user information authentication
	Username string `gorm:"size:30" json:"username"`
	//Password for user information authentication
	Password string `gorm:"size:30" json:"password"`
	//Nickname for user information display
	Nickname string `gorm:"size:50" json:"nickname"`
	Avatar   string `gorm:"size:255" json:"avatar"`
}
