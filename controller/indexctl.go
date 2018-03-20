package controller

import (
	"fmt"
	"net/http"

	"github.com/zidoshare/go-store/confs"
)

func index(w http.ResponseWriter, r *http.Request) {
	currentPage := confs.GetPage(r)
	fmt.Fprintf(w, "hello world and page is %d", currentPage)
}
