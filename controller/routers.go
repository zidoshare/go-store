package controller

import (
	"github.com/gorilla/mux"
)

//Load routers
func Load(router *mux.Router) {
	router.HandleFunc("/", index)
	router.HandleFunc("/api/v0/items", GetItems).Methods("GET")
	router.HandleFunc("/api/v0/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/api/v0/items/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v0/items", AddItem).Methods("POST")
}
