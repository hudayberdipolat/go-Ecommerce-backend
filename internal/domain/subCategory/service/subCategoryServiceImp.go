package service

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	categoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type subcategoryServiceImp struct {
	subCategoryRepo repository.SubCategoryRepository
	categoryRepo    categoryRepository.CategoryRepository
}

func NewSubCategoryService(repo repository.SubCategoryRepository, categoryRepo categoryRepository.CategoryRepository) SubCategoryService {
	return subcategoryServiceImp{
		subCategoryRepo: repo,
		categoryRepo:    categoryRepo,
	}
}

func (subCategoryService subcategoryServiceImp) GetOneSubCategory(categoryID, subCategoryID int) (*dto.GetOneSubCategoryResponse, error) {
	// get category

	category, err := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	subCategory, err := subCategoryService.subCategoryRepo.FindOne(category.ID, subCategoryID)
	if err != nil {
		return nil, err
	}

	subCategoryResponse := dto.NewGetOneSubCategoryResponse(subCategory)
	return &subCategoryResponse, nil
}

func (subCategoryService subcategoryServiceImp) GetAllSubCategory(categoryID int) ([]dto.GetAllSubCategoryResponse, error) {
	// get category
	category, err := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	subCategories, err := subCategoryService.subCategoryRepo.FindAll(category.ID)
	if err != nil {
		return nil, err
	}

	subCategoryResponses := dto.NewGetAllSubCategoryResponse(subCategories)
	return subCategoryResponses, nil
}

func (subCategoryService subcategoryServiceImp) CreateSubCategory(ctx *fiber.Ctx, config *config.Config, categoryID int, createSubCategory dto.CreateSubCategoryRequest) error {

	category, errGetCategory := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if errGetCategory != nil {
		return errors.New("category Not found")
	}
	subCategoryImageURL, err := utils.UploadFile(ctx, "sub_category_image_url", config.FolderConfig.PublicPath, "subCategory-images")
	if err != nil {
		return err
	}
	// create subCategory

	subCategory := models.SubCategory{
		SubCategoryNameTk:   createSubCategory.SubCategoryNameTk,
		SubCategoryNameRu:   createSubCategory.SubCategoryNameRu,
		SubCategoryNameEn:   createSubCategory.SubCategoryNameEn,
		SubCategorySlug:     slug.Make(createSubCategory.SubCategoryNameEn),
		CategoryID:          category.ID,
		SubCategoryImageURL: subCategoryImageURL,
		SubCategoryStatus:   "DRAFT",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	if err := subCategoryService.subCategoryRepo.Store(subCategory); err != nil {
		if err := utils.DeleteFile(*subCategoryImageURL); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (subCategoryService subcategoryServiceImp) UpdateSubCategory(ctx *fiber.Ctx, config *config.Config, categoryID int, subCategoryID int, updateSubCategory dto.UpdateSubCategoryRequest) error {
	category, errGetCategory := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if errGetCategory != nil {
		return errors.New("category Not found")
	}

	subCategory, errSubCategory := subCategoryService.subCategoryRepo.FindOne(category.ID, subCategoryID)
	if errSubCategory != nil {
		return errors.New("subCategory not found")
	}

	file, _ := ctx.FormFile("subCategory_image_url")

	if file != nil {
		// old image delete
		if err := utils.DeleteFile(*subCategory.SubCategoryImageURL); err != nil {
			return err
		}
		// new image upload
		subCategoryImageURL, err := utils.UploadFile(ctx, "sub_category_image_url", config.FolderConfig.PublicPath, "subCategory-images")
		if err != nil {
			return err
		}
		subCategory.SubCategoryImageURL = subCategoryImageURL
	}

	subCategory.SubCategoryNameTk = updateSubCategory.SubCategoryNameTk
	subCategory.SubCategoryNameRu = updateSubCategory.SubCategoryNameRu
	subCategory.SubCategoryNameEn = updateSubCategory.SubCategoryNameEn
	subCategory.SubCategorySlug = slug.Make(updateSubCategory.SubCategoryNameEn)
	subCategory.SubCategoryStatus = updateSubCategory.SubCategoryStatus
	subCategory.UpdatedAt = time.Now()

	if err := subCategoryService.subCategoryRepo.Update(subCategory.ID, *subCategory); err != nil {
		return err
	}
	return nil
}

func (subCategoryService subcategoryServiceImp) DeleteSubCategory(categoryID, subCategoryID int) error {
	category, errGetCategory := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if errGetCategory != nil {
		return errors.New("category Not found")
	}

	subCategory, errSubCategory := subCategoryService.subCategoryRepo.FindOne(category.ID, subCategoryID)
	if errSubCategory != nil {
		return errors.New("subCategory not found")
	}

	if err := subCategoryService.subCategoryRepo.Destroy(category.ID, subCategory.ID); err != nil {
		return nil
	}
	return nil
}
