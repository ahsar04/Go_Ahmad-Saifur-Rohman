package usecase

import (
	"remidial/internal/adapters/repository"
	"remidial/internal/entity"
)

type UserUseCase struct {
	Repo repository.UserRepository
}

func (usecase UserUseCase) GetAllUsers() ([]entity.User, error) {
	users, err := usecase.Repo.GetAllUsers()
	return users, err
}

func (usecase UserUseCase) GetUser(id int) (entity.User, error) {
	user, err := usecase.Repo.GetUser(id)
	return user, err
}

func (usecase UserUseCase) CreateUser(user entity.User) error {
	err := usecase.Repo.CreateUser(user)
	return err
}

func (usecase UserUseCase) UpdateUser(id int, user entity.User) error {
	err := usecase.Repo.UpdateUser(id, user)
	return err
}

func (usecase UserUseCase) DeleteUser(id int) error {
	err := usecase.Repo.DeleteUser(id)
	return err
}

func (usecase UserUseCase) UniqueEmail(email string) error {
	err := usecase.Repo.UniqueEmail(email)
	return err
}
func (usecase UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
	return usecase.Repo.FindByEmail(email)
}

func (usecase UserUseCase) FindUser(id int) error {
	err := usecase.Repo.FindUser(id)
	return err
}
