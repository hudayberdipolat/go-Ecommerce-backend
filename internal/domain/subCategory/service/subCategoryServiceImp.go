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

func (subCategoryService subcategoryServiceImp) GetOneSubCategory(subCategoryID int) (*dto.GetOneSubCategoryResponse, error) {
	subCategory, err := subCategoryService.subCategoryRepo.FindOne(subCategoryID)
	if err != nil {
		return nil, err
	}

	subCategoryResponse := dto.NewGetOneSubCategoryResponse(subCategory)
	return &subCategoryResponse, nil
}

func (subCategoryService subcategoryServiceImp) GetAllSubCategory() ([]dto.GetAllSubCategoryResponse, error) {
	subCategories, err := subCategoryService.subCategoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	subCategoryResponses := dto.NewGetAllSubCategoryResponse(subCategories)
	return subCategoryResponses, nil
}

func (subCategoryService subcategoryServiceImp) CreateSubCategory(ctx *fiber.Ctx, config config.Config, categoryID int, createSubCategory dto.CreateSubCategoryRequest) error {
	// get category eger-de sol category bar bolsa onda subCategory Create edilmeli

	category, errGetCategory := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if errGetCategory != nil {
		return errors.New("category Not found")
	}

	// file update
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
		SubCategoryStatus:   "DRAFT",
		SubCategoryImageURL: subCategoryImageURL,
		CategoryID:          category.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	if err := subCategoryService.subCategoryRepo.Store(subCategory); err != nil {
		return err
	}
	return nil
}

func (subCategoryService subcategoryServiceImp) UpdateSubCategory(ctx *fiber.Ctx, config config.Config, categoryID int, subCategoryID int, updateSubCategory dto.UpdateSubCategoryRequest) error {
	_, errGetCategory := subCategoryService.categoryRepo.FindOneByID(categoryID)
	if errGetCategory != nil {
		return errors.New("category Not found")
	}

	subCategory, errSubCategory := subCategoryService.subCategoryRepo.FindOne(subCategoryID)
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

	subCategory, errSubCategory := subCategoryService.subCategoryRepo.FindOne(subCategoryID)
	if errSubCategory != nil {
		return errors.New("subCategory not found")
	}

	if err := subCategoryService.subCategoryRepo.Destroy(category.ID, subCategory.ID); err != nil {
		return nil
	}
	return nil
}
