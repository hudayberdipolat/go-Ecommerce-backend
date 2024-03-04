package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type CategoryRepository interface {
	// admin ucin functions

	FindOneByID(categoryID int) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Store(category models.Category) error
	Update(categoryID int, category models.Category) error
	Destroy(categoryID int) error

	// user for functions
	FindOneBySlug(categorySlug string) (*models.Category, error)
}
