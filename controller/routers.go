package controller

import (
	"net/http"
	"os"

	"github.com/zidoshare/go-store/logs"

	"github.com/gorilla/mux"
)

//Load routers
func Load() {
	var router = mux.NewRouter()
	var logger = logs.NewLogger(os.Stdout)
	router.HandleFunc("/", index)
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
