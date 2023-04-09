package repository

import "challenge-13/models"

type ProductRepository interface {
	GetProductById(id uint) *models.ProductModel
}