package model

import (
	"github.com/jinzhu/gorm"
)

//Item 商品
type Item struct {
	gorm.Model
	//Title 商品名
	Title string `gorm:"size:100" json:"title"`
	//Price 价格
	Price int `json:"price"`
	//Reserve 余量
	Reserve int `json:"reserve"`
}
