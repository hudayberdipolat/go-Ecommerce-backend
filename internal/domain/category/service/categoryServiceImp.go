package service

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: repo,
	}
}

func (categoryService categoryServiceImp) GetOneCategoryWithID(categoryID int) (*dto.GetOneCategoryResponse, error) {
	category, err := categoryService.categoryRepo.FindOneByID(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewGetOneCategoryResponse(category)
	return &categoryResponse, nil
}

func (categoryService categoryServiceImp) GetAllCategory() ([]dto.GetAllCategoryResponse, error) {
	categories, err := categoryService.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	getAllCategoryResponse := dto.NewGetAllCategoryResponse(categories)
	return getAllCategoryResponse, nil
}

func (categoryService categoryServiceImp) CreateCategory(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateCategoryRequest) error {
	// egerde category_name_[tk, en,ru] bar bolsa onda error gaytarmaly
	// image upload yerine yetmeli dal

	// image upload
	categoryImageURL, err := utils.UploadFile(ctx, "category_image_url", config.FolderConfig.PublicPath, "category-images")
	if err != nil {
		return err
	}

	createCategory := models.Category{
		CategoryNameTk:   strings.ToUpper(createRequest.CategoryNameTk),
		CategoryNameRu:   strings.ToUpper(createRequest.CategoryNameRu),
		CategoryNameEn:   strings.ToUpper(createRequest.CategoryNameEn),
		CategorySlug:     slug.Make(createRequest.CategoryNameEn),
		CategoryStatus:   "DRAFT",
		CategoryImageURL: categoryImageURL,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := categoryService.categoryRepo.Store(createCategory); err != nil {
		return err
	}

	return nil
}

func (categoryService categoryServiceImp) UpdateCategory(ctx *fiber.Ctx, config *config.Config, categoryID int, updateRequest dto.UpdateCategoryRequest) error {

	updateCategory, err := categoryService.categoryRepo.FindOneByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	// image barmy yada yoguny barlamaly we egerde update data-da
	// image bar bolsa onda taze image upload edilmeli we kone image-i delete etmeli
	file, _ := ctx.FormFile("category_image_url")
	if file != nil {
		// old image delete
		if errDelete := utils.DeleteFile(*updateCategory.CategoryImageURL); err != nil {
			return errDelete
		}
		// new image upload
		updateImageURL, err := utils.UploadFile(ctx, "category_image_url", config.FolderConfig.PublicPath, "category-images")
		if err != nil {
			return err
		}
		updateCategory.CategoryImageURL = updateImageURL
	}

	updateCategory.CategoryNameTk = strings.ToUpper(updateRequest.CategoryNameTk)
	updateCategory.CategoryNameRu = strings.ToUpper(updateRequest.CategoryNameRu)
	updateCategory.CategoryNameEn = strings.ToUpper(updateRequest.CategoryNameEn)
	updateCategory.CategorySlug = slug.Make(updateCategory.CategoryNameEn)
	updateCategory.CategoryStatus = updateRequest.CategoryStatus
	updateCategory.UpdatedAt = time.Now()

	if err := categoryService.categoryRepo.Update(updateCategory.ID, *updateCategory); err != nil {
		return err
	}
	return nil
}

func (categoryService categoryServiceImp) DeleteCategory(categoryID int) error {
	deleteCategory, err := categoryService.categoryRepo.FindOneByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	if err := utils.DeleteFile(*deleteCategory.CategoryImageURL); err != nil {
		return err
	}
	if err := categoryService.categoryRepo.Destroy(deleteCategory.ID); err != nil {
		return err
	}
	return nil
}

// front for functions

func (categoryService categoryServiceImp) GetOneCategoryWithSlug(categorySlug string) (*dto.GetOneCategoryResponse, error) {
	panic("category service imp")
}
