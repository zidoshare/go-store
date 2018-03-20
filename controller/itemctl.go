package controller

import (
	"encoding/json"
	"net/http"

	"github.com/zidoshare/go-store/confs"
	"github.com/zidoshare/go-store/model"

	"github.com/zidoshare/go-store/service"
)

//Items get items
func Items(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	current := confs.GetPage(r)
	items, pagination := service.GetItems(current)
	result := confs.Success(&struct{
		items
	})
	json.NewEncoder(w).Encode(result)
}

//AddItem add item
func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := &model.Item{}
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = nil
	item.CreatedAt = nil
	item.DeletedAt = nil
	item.UpdatedAt = nil
	if item.Title == nil || item.Title == "" {
		json.NewEncoder(w).Encode(confs.Fail())
		return
	}
	json.NewEncoder(w).Encode(confs.Success(nil))
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header.Set("Content-Type", "application/json")
	r.ParseForm()

}
