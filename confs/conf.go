package confs

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/zidoshare/go-store/logs"
)

var logger = logs.NewLogger(os.Stdout)

// Conf go-store
var Conf *Configuration

// Configuration of go-store
type Configuration struct {
	Mysql       string
	LogLevel    string
	RuntimeMode string
	Wait        time.Duration
	PageSize    int
	Server      string
}

// Load load store base configuration
func Load() {
	mysql := flag.String("mysql", "", "the mysql server like \"username:password@(localhost:3306)/store?charset=utf8mb4&parseTime=True&loc=Local\"")
	logLevel := flag.String("log_level", "", "logging level: trace/debug/info/warn/error/fatal")
	runtimeMode := flag.String("mode", "", "runtime mode (dev/prod)")
	path := flag.String("path", "conf.json", "the config path")
	pageSize := flag.Int("page_size", 0, "the page size")
	listen := flag.String("listen", "", "listening server like \":8080\"")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
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
	Conf.Wait = wait
	if *pageSize != 0 {
		Conf.PageSize = *pageSize
	}

	if *listen != "" {
		Conf.Server = *listen
	}
	logs.SetLevel(Conf.LogLevel)
	logger.Debugf("config end... confg:%+v", Conf)
}
