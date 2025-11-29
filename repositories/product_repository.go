package repositories

import "testMKP/models"

type ProductRepository interface {
	Create(product models.ProductCreateRequest) (models.Product, error)
	Update(product models.ProductUpdateRequest) (models.Product, error)
	Delete(id string) error
	FindAll() ([]models.Product, error)
	FindById(id string) (models.Product, error)
}
