package service

import (
	"github.com/zidoshare/go-store/confs"
	"github.com/zidoshare/go-store/model"
)

//GetItems get item list
func GetItems(current int) (items []*model.Item, pagination *confs.Pagination) {
	offset := (current - 1) * pageSize
	total := 0
	if err := db.Count(&total).Limit(pageSize).Offset(offset).Find(&items).Error; err != nil {
		logger.Errorf("get items failed: " + err.Error())
	}
	pagination = &confs.Pagination{
		Current:  current,
		PageSize: pageSize,
		Total:    total,
	}
	return
}

//AddItem add item
func AddItem(item *model.Item) error {
	return db.Create(&item).Error
}

//DeleteItem delete item by id
func DeleteItem(id uint) error {
	return db.Where("id = ?", id).Delete(&model.Item{}).Error
}
