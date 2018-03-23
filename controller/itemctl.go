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
	w.Header().Set("Content-Type", "application/json")
	current := common.GetPage(r)
	items, pagination := service.GetItems(current)
	result := &common.Resp{
		Code:       0,
		Data:       items,
		Pagination: pagination,
	}
	json.NewEncoder(w).Encode(result)
}

//GetItem get item by id
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	item := service.GetItem(uint(id))
	json.NewEncoder(w).Encode(&common.Resp{
		Code: 0,
		Data: item,
	})
}

//AddItem add item
func AddItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close
	w.Header().Set("Content-Type", "application/json")
	item := &model.Item{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	err = json.Unmarshal(bytes, item)
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	err = service.AddItem(item)
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	if item.Title == "" {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	json.NewEncoder(w).Encode(common.Resp{
		Code: 0,
	})
}

//DeleteItem delete item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	err = service.DeleteItem(uint(id))
	if err != nil {
		json.NewEncoder(w).Encode(common.Resp{
			Code: -1,
		})
		return
	}
	json.NewEncoder(w).Encode(common.Resp{
		Code: 0,
	})
}
