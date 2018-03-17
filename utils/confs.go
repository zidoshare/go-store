package utils

import (
	"flag"
	"io/ioutil"
	"os"

	"encoding/json"

	"github.com/zidoshare/go-store/logs"
)

var logger = logs.NewLogger(os.Stdout)

//Conf go-store
var Conf *Configuration

//Configuration of go-store
type Configuration struct {
	Mysql       string
	LogLevel    string
	RuntimeMode string
}

//Load load store base configuration
func Load() {
	mysql := flag.String("mysql", "", "the mysql server like \"username:password@(localhost:3306)/stroe?charset=utf8mb4&parseTime=True&loc=Local\"")
	logLevel := flag.String("log_level", "", "logging level: trace/debug/info/warn/error/fatal")
	runtimeMode := flag.String("mode", "", "runtime mode (dev/prod)")
	path := flag.String("path", "conf.json", "the config path")
	flag.Parse()
	bytes, err := ioutil.ReadFile(*path)
	if err != nil {
		logger.Fatal("loads configuration file [" + *path + "] failed: " + err.Error())
	}
	Conf = new(Configuration)
	if err = json.Unmarshal(bytes, Conf); err != nil {
		logger.Fatal("parse json configuration failed", err)
	}
	if *logLevel != "" {
		Conf.LogLevel = *logLevel
	}
	if *mysql != "" {
		Conf.Mysql = *mysql
	}
	if *runtimeMode != "" {
		Conf.RuntimeMode = *runtimeMode
	}

	logs.SetLevel(Conf.LogLevel)
}
