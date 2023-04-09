package repository

import (
	"challenge-13/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) GetProductById(id string) *models.ProductModel {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	}

	product := args.Get(0).(models.ProductModel)

	return &product
}