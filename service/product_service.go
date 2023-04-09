package service

import (
	"challenge-13/models"
	"challenge-13/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id uint) (*models.ProductModel, error) {
	product := service.Repository.GetProductById(id)
	if product == nil {
		return nil, errors.New("Product not found!")
	}

	return product, nil
}