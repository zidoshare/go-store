package service

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zidoshare/go-store/logs"
)

var logger = logs.NewLogger(os.Stdout)
var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(localhost:3306)/stroe?charset=utf8mb4&parseTime=True&loc=Local")
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
}
