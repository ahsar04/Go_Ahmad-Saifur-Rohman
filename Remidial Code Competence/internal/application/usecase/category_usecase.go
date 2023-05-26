package usecase

import (
	"remidial/internal/adapters/repository"
	"remidial/internal/entity"
)

type CategoryUseCase struct {
	Repo repository.CategoryRepository
}

func (usecase CategoryUseCase) GetAllCategories() ([]entity.Category, error) {
	Categories, err := usecase.Repo.GetAllCategories()
	return Categories, err
}

func (usecase CategoryUseCase) GetCategory(id int) (entity.Category, error) {
	Category, err := usecase.Repo.GetCategory(id)
	return Category, err
}

func (usecase CategoryUseCase) CreateCategory(Category entity.Category) error {
	err := usecase.Repo.CreateCategory(Category)
	return err
}

func (usecase CategoryUseCase) UpdateCategory(id int, Category entity.Category) error {
	err := usecase.Repo.UpdateCategory(id, Category)
	return err
}
func (usecase CategoryUseCase) FindCategory(id int) error {
	err := usecase.Repo.FindCategory(id)
	return err
}

func (usecase CategoryUseCase) DeleteCategory(id int) error {
	err := usecase.Repo.DeleteCategory(id)
	return err
}

