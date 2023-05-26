package repository

import (
	"remidial/internal/entity"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (repo ItemRepository) GetAllItems() ([]entity.Item, error) {
	var items []entity.Item
	result := repo.DB.Preload("Category").Find(&items)
	return items, result.Error
}

func (repo ItemRepository) GetItem(id int) (entity.Item, error) {
	var items entity.Item
	result := repo.DB.Preload("Category").First(&items, id)
	return items, result.Error
}

func (repo ItemRepository) GetItemByCategory(id_category int) (entity.Item, error) {
	var items entity.Item
	result := repo.DB.Preload("Category").First(&items, "category_id = ?",id_category)
	return items, result.Error
}
func (repo ItemRepository) GetItemByName(name string) (entity.Item, error) {
	var items entity.Item
	result := repo.DB.Preload("Category").Find(&items, "item_name like ?","%"+name+"%")
	return items, result.Error
}

func (repo ItemRepository) CreateItem(item entity.Item) error {
	result := repo.DB.Preload("Category").Create(&item)
	return result.Error
}

func (repo ItemRepository) UpdateItem(id int, item entity.Item) error {
	result := repo.DB.Preload("Category").Model(&item).Where("id = ?", id).Updates(&item)
	return result.Error
}

func (repo ItemRepository) DeleteItem(id int) error {
	result := repo.DB.Delete(&entity.Item{}, id)
	return result.Error
}

func (repo ItemRepository) FindItem(id int) error {
	result := repo.DB.Preload("Category").First(&entity.Item{}, id)
	return result.Error
}
