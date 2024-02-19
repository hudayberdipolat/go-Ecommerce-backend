package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: repo,
	}
}

func (service categoryServiceImp) FindOneCategory(categoryID int) (*dto.OneCategoryResponse, error) {
	panic("category service imp")
}

func (service categoryServiceImp) FindAllCategory() ([]dto.AllCategoryResponse, error) {
	panic("category service imp")
}

func (service categoryServiceImp) CreateCategory(request dto.CreateCategoryRequest) error {
	panic("category service imp")
}

func (service categoryServiceImp) UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error {
	panic("category service imp")
}

func (service categoryServiceImp) DeleteCategory(categoryID int) error {
	panic("category service imp")
}
