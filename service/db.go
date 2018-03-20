package service

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql support
	"github.com/zidoshare/go-store/confs"
	"github.com/zidoshare/go-store/logs"
	"github.com/zidoshare/go-store/model"
)

var logger = logs.NewLogger(os.Stdout)
var db *gorm.DB
var pageSize = confs.Conf.PageSize

//Connect connect store database
func Connect() {
	var err error
	db, err = gorm.Open("mysql", confs.Conf.Mysql)
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	if err = db.AutoMigrate(&model.Item{}, &model.User{}).Error; err != nil {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

//DisConnect desconnect store database
func DisConnect() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}
