package controller

import (
	"github.com/gorilla/mux"
)

//Load routers
func Load(router *mux.Router) {
	router.HandleFunc("/", index)
	router.HandleFunc("/api/v0/items", Items).Methods("GET")
}
