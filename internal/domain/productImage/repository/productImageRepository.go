package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductImageRepository interface {
	GetOne(productImageID, productID int) (*models.ProductImage, error)
	Create(image models.ProductImage) error
	Delete(productImageID, productID int) error
	GetOneProduct(productID int) (*models.Product, error)
}
