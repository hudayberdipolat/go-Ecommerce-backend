package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductImageRepository interface {
	GetOneProduct(productID int) (*models.Product, error)
	GetOne(productImageID int) (*models.ProductImage, error)
	Create(image models.ProductImage) error
	Delete(productImageID int) error
}
