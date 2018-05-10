package controller

import (
	"github.com/gorilla/mux"
	"github.com/zidoshare/go-store/logs"
)

//StoreRouters init by routers
var StoreRouters = mux.NewRouter()

var logger = logs.NewStdLogger()

func init() {
	//load routers
	logger.Info("start loading routers")
	StoreRouters.HandleFunc("/", index)
	StoreRouters.HandleFunc("/api/v0/items", GetItems).Methods("GET")
	StoreRouters.HandleFunc("/api/v0/items/{id}", GetItem).Methods("GET")
	StoreRouters.HandleFunc("/api/v0/items/{id}", DeleteItem).Methods("DELETE")
	StoreRouters.HandleFunc("/api/v0/items", AddItem).Methods("POST")
}
