package service

import (
	"challenge-13/models"
	"challenge-13/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", 1).Return(nil)
	product, err := productService.GetOneProduct(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found!", err.Error(), "Error response has to be 'Product not found!'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.ProductModel{
		Title: "Kaca",
	}
	product.ID = 1

	productRepository.Mock.On("GetProductById", 1).Return(product)

	result, err := productService.GetOneProduct(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.ID, result.ID, "Result has to be '1'")
	assert.Equal(t, product.Title, result.Title, "Result has to be 'Kaca'")
	assert.Equal(t, &product, result, "Result has to be a product data with id '1'")
}