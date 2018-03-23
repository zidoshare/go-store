package common

import (
	"encoding/json"
	"net/http"
)

//Code reponse code value
type Code int

const (
	//SUCCESS code
	SUCCESS Code = 0
)

//Resp the REST response model
type Resp struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
	Err        *RespError  `json:"err"`
}

//RespondWithJSON write json and code to response
func RespondWithJSON(w http.ResponseWriter, code int, resp *Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func RespondWithError(w http.ResponseWriter, code int, msg *RespError) {
	RespondWithJSON(w, code, &Resp{})
}
