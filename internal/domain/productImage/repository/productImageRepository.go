package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductImageRepository interface {
	GetOneProductImage(productID, productImageID int) (*models.ProductImage, error)
	Create(image models.ProductImage) error
	Delete(productID, productImageID int) error
	GetOneProduct(productID int) (*models.Product, error)
}
