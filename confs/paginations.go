package confs

import (
	"net/http"
	"strconv"
)

// Pagination this page model
type Pagination struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

// GetPage get page from http.Request
func GetPage(r *http.Request) int {
	vals := r.URL.Query()
	param := vals["p"]
	ret, _ := strconv.Atoi(param[0])
	if 1 > ret {
		ret = 1
	}
	return ret
}
