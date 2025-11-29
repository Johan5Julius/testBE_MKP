package services

import (
	"errors"
	"testMKP/helper"
	"testMKP/models"
	"testMKP/repositories"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepo}
}

func (p *ProductServiceImpl) Create(product models.ProductCreateRequest) (models.Product, error) {
	if err := helper.ValidateRequired("name", product.Name); err != nil {
		return models.Product{}, err
	}

	if err := helper.ValidateLength("name", product.Name, 3, 100); err != nil {
		return models.Product{}, err
	}

	if product.Price <= 0 {
		return models.Product{}, errors.New("price harus lebih besar dari 0")
	}

	return p.ProductRepository.Create(product)
}

func (p *ProductServiceImpl) Update(product models.ProductUpdateRequest) (models.Product, error) {
	if product.ID == "" {
		return models.Product{}, errors.New("ID tidak boleh kosong")
	}

	if err := helper.ValidateRequired("name", product.Name); err != nil {
		return models.Product{}, err
	}

	if err := helper.ValidateLength("name", product.Name, 3, 100); err != nil {
		return models.Product{}, err
	}

	if product.Price <= 0 {
		return models.Product{}, errors.New("price harus lebih besar dari 0")
	}

	return p.ProductRepository.Update(product)
}

func (p *ProductServiceImpl) Delete(id string) error {
	if id == "" {
		return errors.New("ID tidak boleh kosong")
	}
	return p.ProductRepository.Delete(id)
}

func (p *ProductServiceImpl) FindById(id string) (models.Product, error) {
	if id == "" {
		return models.Product{}, errors.New("ID tidak boleh kosong")
	}
	return p.ProductRepository.FindById(id)
}

func (p *ProductServiceImpl) FindAll() ([]models.Product, error) {
	return p.ProductRepository.FindAll()
}
