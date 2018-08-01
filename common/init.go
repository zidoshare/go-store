//Package common provide some common support
package common

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/zidoshare/go-store/logs"
)

var logger = logs.NewLogger(os.Stdout)
var _testing = false

// init load store base configuration
func init() {
	logs.SetLevel("info")
	if _testing {
		logger.Info("testing,skip init")
		return
	}
	logger.Info("load store base configurations")
	mysql := flag.String("mysql", "", "the mysql server like \"username:password@(localhost:3306)/store?charset=utf8mb4&parseTime=True&loc=Local\"")
	logLevel := flag.String("log_level", "", "logging level: trace/debug/info/warn/error/fatal")
	runtimeMode := flag.String("mode", "", "runtime mode (dev/prod)")
	path := flag.String("path", "conf.json", "the config path")
	pageSize := flag.Int("page_size", 0, "the page size")
	listen := flag.String("listen", "", "listening server like \":8080\"")
	iss := flag.String("iss", "", "set token iss")
	spwd := flag.String("spwd", "", "secure password")
	loginExp := flag.Int("login_exp", 0, "login expiration time (s)")
	alg := flag.String("alg", "", "secure method split by ','")
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

	if *iss != "" {
		Conf.Iss = *iss
	}

	if *loginExp != 0 {
		Conf.LoginExp = *loginExp
	}

	if *spwd != "" {
		Conf.Spwd = *spwd
	}
	if *alg != "" {
		Conf.Alg = strings.Split(*alg, ",")
	}
	logs.SetLevel(Conf.LogLevel)
	logger.Debugf("current log level -> %s", Conf.LogLevel)
	logger.Debugf("current mysql connection -> %s,", Conf.Mysql)
	logger.Debugf("current runtime mode -> %s", Conf.RuntimeMode)
	logger.Debugf("current pageSize -> %d", Conf.PageSize)
	logger.Debugf("current server -> %s", Conf.Server)
	logger.Debugf("current iss -> %s", Conf.Iss)
	logger.Debugf("current loginExp -> %d", Conf.LoginExp)
	logger.Debugf("current alg -> %s", strings.Join(Conf.Alg, ","))
	logger.Debugf("current spwd -> %s", Conf.Spwd)
}
