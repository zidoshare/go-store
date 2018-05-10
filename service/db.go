package service

import (
	"os"

	_ "github.com/go-sql-driver/mysql" //mysql support
	"github.com/jinzhu/gorm"
	"github.com/zidoshare/go-store/common"
	"github.com/zidoshare/go-store/logs"
	"github.com/zidoshare/go-store/model"
)

var logger = logs.NewLogger(os.Stdout)
var db *gorm.DB

//Connect connect store database
func Connect() {
	logger.Infof("connect to mysql,url:%s", common.Conf.Mysql)
	var err error
	db, err = gorm.Open("mysql", common.Conf.Mysql)
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	if err = db.AutoMigrate(&model.Item{}, &model.User{}).Error; err != nil {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}
	db.LogMode(common.Dev())
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

//DisConnect desconnect store database
func DisConnect() {
	logger.Info("disconnect from database")
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}
