package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type categoryRepositoryImp struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepositoryImp{
		db: db,
	}
}

func (repo categoryRepositoryImp) OneCategory(categoryID int) (*models.Category, error) {
	panic("category repository imp")
}

func (repo categoryRepositoryImp) AllCategory() ([]models.Category, error) {
	panic("category repository imp")
}

func (repo categoryRepositoryImp) Create(category models.Category) error {
	panic("category repository imp")
}

func (repo categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	panic("category repository imp")
}

func (repo categoryRepositoryImp) Delete(categoryID int) error {
	panic("category repository imp")
}
