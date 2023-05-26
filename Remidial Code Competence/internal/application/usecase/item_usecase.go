package usecase

import (
	"remidial/internal/adapters/repository"
	"remidial/internal/entity"
)

type ItemUseCase struct {
	Repo repository.ItemRepository
}

func (usecase ItemUseCase) GetAllItems() ([]entity.Item, error) {
	Items, err := usecase.Repo.GetAllItems()
	return Items, err
}

func (usecase ItemUseCase) GetItem(id int) (entity.Item, error) {
	Item, err := usecase.Repo.GetItem(id)
	return Item, err
}
func (usecase ItemUseCase) GetItemByCategory(id_category int) (entity.Item, error) {
	Item, err := usecase.Repo.GetItemByCategory(id_category)
	return Item, err
}
func (usecase ItemUseCase) GetItemByName(name string) (entity.Item, error) {
	Item, err := usecase.Repo.GetItemByName(name)
	return Item, err
}

func (usecase ItemUseCase) CreateItem(Item entity.Item) error {
	err := usecase.Repo.CreateItem(Item)
	return err
}

func (usecase ItemUseCase) UpdateItem(id int, Item entity.Item) error {
	err := usecase.Repo.UpdateItem(id, Item)
	return err
}

func (usecase ItemUseCase) DeleteItem(id int) error {
	err := usecase.Repo.DeleteItem(id)
	return err
}

func (usecase ItemUseCase) FindItem(id int) error {
	err := usecase.Repo.FindItem(id)
	return err
}
