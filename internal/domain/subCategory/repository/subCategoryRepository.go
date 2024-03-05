package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type SubCategoryRepository interface {
	FindOne(subCategoryID int) (*models.SubCategory, error)
	FindAll() ([]models.SubCategory, error)
	Store(createRequest models.SubCategory) error
	Update(subCategoryID int, updateRequest models.SubCategory) error
	Destroy(categoryID, subCategoryID int) error

	// for check

}
