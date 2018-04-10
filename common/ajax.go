package common

import (
	"encoding/json"
	"net/http"
)

//Code reponse code value
type Code int

const (
	//ErrorParam get param error
	ErrorParam = 1001
	//ErrorParseRequest parse request params error
	ErrorParseRequest = 1002
	//ErrorParamInValid param is invalid
	ErrorParamInValid = 1003
	//ErrorResolveRequest resolve request error
	ErrorResolveRequest = 1005
	//ErrorTokenInvalid token is invalid
	ErrorTokenInvalid = 1
)

//Resp the REST response model
type Resp struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

//RespondJSON write json to response
func RespondJSON(w http.ResponseWriter, status int, resp *Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if resp != nil {
		json.NewEncoder(w).Encode(resp)
	}
}

//RespondError write Error and code to response
func RespondError(w http.ResponseWriter, status int, respErr *RespErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if respErr != nil {
		json.NewEncoder(w).Encode(respErr)
	}
}

// RespondBadRequestError can usually be returned when the request parameter is invalid or malformed.
func RespondBadRequestError(w http.ResponseWriter, err *RespErr) {
	RespondError(w, 400, err)
}

// RespondUnauthorizedError The requested resource requires authentication.
// The client does not provide authentication information or the authentication information is incorrect.
func RespondUnauthorizedError(w http.ResponseWriter) {
	RespondError(w, 401, nil)
}

// RespondForbidden This status code is used when there is no permission to operate the resource
// (such as modify/delete a resource that does not belong to the user)
func RespondForbidden(w http.ResponseWriter) {
	RespondError(w, 403, nil)
}

//RespondNotFound respond 404 status ,can not find the resource
func RespondNotFound(w http.ResponseWriter) {
	RespondError(w, 404, nil)
}

//RespondData response 200 status ,find the resource successfuly
func RespondData(w http.ResponseWriter, data interface{}) {
	RespondJSON(w, 200, &Resp{
		Data: data,
	})
}

//RespondPage response 200 status,find a page
func RespondPage(w http.ResponseWriter, data interface{}, pagination *Pagination) {
	RespondJSON(w, 200, &Resp{
		Data:       data,
		Pagination: pagination,
	})
}

//RespondStatus don't need respond anything
func RespondStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

//RespondCreateOk status 201. Resources are created correctly. eg:create
func RespondCreateOk(w http.ResponseWriter, location string) {
	RespondStatus(w, 201)
	w.Header().Set("Location", location)
}

//RespondNoData status 204 ，The request is correct but there is no content to return . eg:delete/update
func RespondNoData(w http.ResponseWriter) {
	RespondStatus(w, 204)
}
