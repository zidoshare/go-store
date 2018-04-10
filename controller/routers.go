package controller

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zidoshare/go-store/common"
)

func handleToken(w http.ResponseWriter, r *http.Request, exceptedRole string) (uint, bool) {
	uid, role, err := common.ParseToken(w, r)
	if err != nil {
		common.RespondUnauthorizedError(w)
		return 0, false
	}
	if strings.Contains(exceptedRole, role) {
		common.RespondForbidden(w)
		return 0, false
	}
	return uid, true
}

//Admin Determine if it is admin
// return uid(uint)
func admin(w http.ResponseWriter, r *http.Request) uint {
	uid, _ := handleToken(w, r, "admin")
	return uid
}

//checkAdmin Determine if it is admin
// return true/false
func checkAdmin(w http.ResponseWriter, r *http.Request) bool {
	_, ok := handleToken(w, r, "admin")
	return ok
}

//user Determine if it is admin/user
// return uid(uint)
func user(w http.ResponseWriter, r *http.Request) uint {
	uid, _ := handleToken(w, r, "admin|user")
	return uid
}

//checkUser Determine if it is admin/user
// return true/false
func checkUser(w http.ResponseWriter, r *http.Request) bool {
	_, ok := handleToken(w, r, "admin|user")
	return ok
}

//Load routers
func Load(router *mux.Router) {
	router.HandleFunc("/", index)
	router.HandleFunc("/api/v0/items", GetItems).Methods("GET")
	router.HandleFunc("/api/v0/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/api/v0/items/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v0/items", AddItem).Methods("POST")
}
