package repositories

import "testMKP/models"

type UserRepository interface {
	Create(user models.UserCreateRequest) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Update(user models.UserUpdateRequest) (models.User, error)
	Delete(id int) error
}
