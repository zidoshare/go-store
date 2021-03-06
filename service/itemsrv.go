package service

import (
	"github.com/zidoshare/go-store/common"
	"github.com/zidoshare/go-store/model"
)

//GetItems get item list
func GetItems(current int) (items []*model.Item, pagination *common.Pagination) {
	pageSize := common.Conf.PageSize
	offset := (current - 1) * pageSize
	total := 0
	if err := db.Model(&model.Item{}).Count(&total).Limit(pageSize).Offset(offset).Find(&items).Error; err != nil {
		logger.Errorf("get items failed: " + err.Error())
	}
	pagination = &common.Pagination{
		Current:  current,
		PageSize: pageSize,
		Total:    total,
	}
	return
}

//GetItem get item by primary key
func GetItem(id uint) *model.Item {
	item := &model.Item{}
	db.Find(item, id)
	return item
}

//AddItem add item
func AddItem(item *model.Item) error {
	return db.Create(&item).Error
}

//DeleteItem delete item by id
func DeleteItem(id uint) error {
	return db.Where("id = ?", id).Delete(&model.Item{}).Error
}
