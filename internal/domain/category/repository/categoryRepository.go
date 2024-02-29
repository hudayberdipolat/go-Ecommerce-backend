package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type CategoryRepository interface {
	OneCategory(categoryID int) (*models.Category, error)
	AllCategory() ([]models.Category, error)
	Create(category models.Category) error
	Update(categoryID int, category models.Category) error
	Delete(categoryID int) error
	GetOneCategoryWithSlug(categorySlug string) (*models.Category, error)
}
