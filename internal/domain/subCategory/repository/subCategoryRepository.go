package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type SubCategoryRepository interface {
	FindOne(categoryID, subCategoryID int) (*models.SubCategory, error)
	FindAll(categoryID int) ([]models.SubCategory, error)
	Store(createRequest models.SubCategory) error
	Update(subCategoryID int, updateRequest models.SubCategory) error
	Destroy(categoryID, subCategoryID int) error

	// for check

}
