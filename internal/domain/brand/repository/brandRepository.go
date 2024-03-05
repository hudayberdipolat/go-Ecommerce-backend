package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type BrandRepository interface {
	GetOneByID(brandID int) (*models.Brand, error)
	GetAll() ([]models.Brand, error)
	Store(brand models.Brand) error
	Update(brandID int, brand models.Brand) error
	Destroy(brandID int) error
}
