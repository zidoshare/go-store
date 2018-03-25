package common

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
func GetPage(r *http.Request) (int, error) {
	vals := r.URL.Query()
	param := vals["p"]
	if param == nil {
		return 1, nil
	}
	ret, err := strconv.Atoi(param[0])
	if 1 > ret {
		ret = 1
	}
	return ret, err
}
