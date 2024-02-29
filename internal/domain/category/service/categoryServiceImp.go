package service

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
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
	category, err := service.categoryRepo.OneCategory(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewOneCategoryResponse(category)
	return &categoryResponse, nil
}

func (service categoryServiceImp) FindAllCategory() ([]dto.AllCategoryResponse, error) {
	categories, err := service.categoryRepo.AllCategory()
	if err != nil {
		return nil, err
	}
	allCategoryResponse := dto.NewAllCategoryResponse(categories)
	return allCategoryResponse, nil
}

func (service categoryServiceImp) CreateCategory(request dto.CreateCategoryRequest) error {

	createCategory := models.Category{
		CategoryNameTK: request.CategoryNameTK,
		CategoryNameRU: request.CategoryNameRU,
		CategoryStatus: request.CategoryStatus,
		CategorySlug:   slug.Make(request.CategoryNameTK),
	}
	if err := service.categoryRepo.Create(createCategory); err != nil {
		return err
	}
	return nil
}

func (service categoryServiceImp) UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error {

	// get category
	getCategory, err := service.categoryRepo.OneCategory(categoryID)
	if err != nil {
		return errors.New("category not found! something wrong ...")
	}

	getCategory.CategoryNameTK = request.CategoryNameTK
	getCategory.CategoryNameRU = request.CategoryNameRU
	getCategory.CategoryStatus = request.CategoryStatus
	getCategory.CategorySlug = slug.Make(getCategory.CategoryNameTK)

	// update category

	if err := service.categoryRepo.Update(getCategory.ID, *getCategory); err != nil {
		return err
	}
	return nil
}

func (service categoryServiceImp) DeleteCategory(categoryID int) error {
	getCategory, err := service.categoryRepo.OneCategory(categoryID)
	if err != nil {
		return errors.New("category not found! something wrong ...")
	}

	if err := service.categoryRepo.Delete(getCategory.ID); err != nil {
		return err
	}
	return nil
}

func (service categoryServiceImp) GetOneCategory(categorySlug string) (*dto.OneCategoryResponse, error) {
	getCategory, err := service.categoryRepo.GetOneCategoryWithSlug(categorySlug)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewOneCategoryResponse(getCategory)

	return &categoryResponse, nil
}
