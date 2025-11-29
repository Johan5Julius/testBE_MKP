package services

import "testMKP/models"

type ProductService interface {
	Create(product models.ProductCreateRequest) (models.Product, error)
	Update(product models.ProductUpdateRequest) (models.Product, error)
	Delete(id string) error
	FindById(id string) (models.Product, error)
	FindAll() ([]models.Product, error)
}
