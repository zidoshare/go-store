package common

import (
	"time"
)

// Conf go-store
var Conf *Configuration

// Configuration of go-store
type Configuration struct {
	Mysql       string        //mysql connection message
	LogLevel    string        //log level
	RuntimeMode string        //runtime mode dev/prod
	Wait        time.Duration //waiting on close server
	PageSize    int           //page size
	Server      string        //listen server
	Iss         string        //issuer
	LoginExp    int           //login expiration time (s)
	Spwd        string        //hs256 password
	Alg         []string      //array of secure method
}

//Dev return true if current runtime mode is dev
func Dev() bool {
	return Conf.RuntimeMode == "dev"
}
