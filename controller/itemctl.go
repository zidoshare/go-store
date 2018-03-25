package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/zidoshare/go-store/common"
	"github.com/zidoshare/go-store/model"

	"github.com/zidoshare/go-store/service"
)

//GetItems get items
func GetItems(w http.ResponseWriter, r *http.Request) {
	current, err := common.GetPage(r)
	if err != nil {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:       common.ErrorParam,
			Message:    "this param p is bad",
			Additional: "p",
		})
		return
	}
	items, pagination := service.GetItems(current)
	common.RespondPage(w, items, pagination)
}

//GetItem get item by id
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:       common.ErrorParam,
			Message:    "this param id is bad",
			Additional: "id",
		})
		return
	}
	item := service.GetItem(uint(id))
	common.RespondJSON(w, http.StatusOK, &common.Resp{
		Data: item,
	})
}

//AddItem add item
func AddItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	item := &model.Item{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:    common.ErrorParseRequest,
			Message: err.Error(),
		})
		return
	}
	err = json.Unmarshal(bytes, item)
	if err != nil {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:    common.ErrorParseRequest,
			Message: err.Error(),
		})
		return
	}
	if item.Title == "" {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:       common.ErrorParamInValid,
			Message:    "param is invalid",
			Additional: "title",
		})
		return
	}
	err = service.AddItem(item)
	if err != nil {
		common.RespondError(w, 500, &common.RespErr{
			Code:    common.ErrorResolveRequest,
			Message: err.Error(),
		})
		return
	}

	common.RespondCreateOk(w, strconv.Itoa(int(item.ID)))
}

//DeleteItem delete item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.RespondBadRequestError(w, &common.RespErr{
			Code:       common.ErrorParam,
			Message:    "this param id is bad",
			Additional: "id",
		})
		return
	}
	err = service.DeleteItem(uint(id))
	if err != nil {
		common.RespondError(w, 500, &common.RespErr{
			Code:    common.ErrorResolveRequest,
			Message: err.Error(),
		})
		return
	}
	common.RespondNoData(w)
}
