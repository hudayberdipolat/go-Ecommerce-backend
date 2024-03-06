package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductImageRepository interface {
	FindAll() ([]models.ProductImage, error)
	FindOne(productID, productImageID int) (*models.ProductImage, error)
	Store(productImage models.ProductImage) error
	Destroy(productID, productImageID int) error
}
