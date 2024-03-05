package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductRepository interface {
	GetOneByID(productID int) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Store(product models.Product) error
	Update(productID int, product models.Product) error
	Destroy(productID int) error
}
