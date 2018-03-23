package controller

import (
	"fmt"
	"net/http"

	"github.com/zidoshare/go-store/common"
)

func index(w http.ResponseWriter, r *http.Request) {
	currentPage := common.GetPage(r)
	fmt.Fprintf(w, "hello world and page is %d", currentPage)
}
