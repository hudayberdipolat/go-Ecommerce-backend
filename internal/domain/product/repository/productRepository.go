package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductRepository interface {
	FindOne(productID int) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Create(product models.Product) error
	Update(productID int, product models.Product) error
	Delete(productID int) error

	FindOneProductWithSlug(productSlug string) (*models.Product, error)
}
