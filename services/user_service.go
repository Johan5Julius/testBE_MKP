package services

import "testMKP/models"

type UserService interface {
	Create(user models.UserCreateRequest) (models.User, error)
	Update(user models.UserUpdateRequest) (models.User, error)
	Delete(id int) error
	FindById(id int) (models.User, error)
	FindAll() ([]models.User, error)
}
