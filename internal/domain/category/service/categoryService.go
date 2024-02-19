package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"

type CategoryService interface {
	FindOneCategory(categoryID int) (*dto.OneCategoryResponse, error)
	FindAllCategory() ([]dto.AllCategoryResponse, error)
	CreateCategory(request dto.CreateCategoryRequest) error
	UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}
