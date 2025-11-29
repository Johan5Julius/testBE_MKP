package services

import (
	"testMKP/helper"
	"testMKP/models"
	"testMKP/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepo}
}

func (u UserServiceImpl) Create(user models.UserCreateRequest) (models.User, error) {
	if err := helper.ValidateRequired("name", user.Name); err != nil {
		return models.User{}, err
	}

	if err := helper.ValidateLength("name", user.Name, 3, 100); err != nil {
		return models.User{}, err
	}

	if err := helper.ValidateEmail(user.Email); err != nil {
		return models.User{}, err
	}

	return u.UserRepository.Create(user)
}

func (u UserServiceImpl) Update(user models.UserUpdateRequest) (models.User, error) {
	if err := helper.ValidateRequired("name", user.Name); err != nil {
		return models.User{}, err
	}

	if err := helper.ValidateLength("name", user.Name, 3, 100); err != nil {
		return models.User{}, err
	}

	if err := helper.ValidateEmail(user.Email); err != nil {
		return models.User{}, err
	}

	return u.UserRepository.Update(user)
}

func (u UserServiceImpl) Delete(id int) error {
	return u.UserRepository.Delete(id)
}

func (u UserServiceImpl) FindById(id int) (models.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u UserServiceImpl) FindAll() ([]models.User, error) {
	return u.UserRepository.FindAll()
}
