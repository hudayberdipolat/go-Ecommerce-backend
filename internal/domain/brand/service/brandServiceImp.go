package service

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type brandServiceImp struct {
	brandRepo repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) BrandService {
	return brandServiceImp{
		brandRepo: repo,
	}
}

func (brandService brandServiceImp) GetOneBrandByID(brandID int) (*dto.GetOneBrandResponse, error) {
	brand, err := brandService.brandRepo.GetOneByID(brandID)
	if err != nil {
		return nil, err
	}
	brandResponse := dto.NewGetOneBrandResponse(brand)
	return &brandResponse, nil
}

func (brandService brandServiceImp) GetAllBrand() ([]dto.GetAllBrandResponse, error) {
	brands, err := brandService.brandRepo.GetAll()
	if err != nil {
		return nil, err
	}
	brandResponses := dto.NewGetAllBrandResponse(brands)
	return brandResponses, nil
}

func (brandService brandServiceImp) CreateBrand(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateBrandRequest) error {

	// file upload
	brandImageURL, err := utils.UploadFile(ctx, "brand_image_url", config.FolderConfig.PublicPath, "brand-images")
	if err != nil {
		return err
	}

	// create brand
	createBrand := models.Brand{
		BrandNameTk:   createRequest.BrandNameTk,
		BrandNameRu:   createRequest.BrandNameRu,
		BrandNameEn:   createRequest.BrandNameEn,
		BrandImageURL: brandImageURL,
		BrandSlug:     slug.Make(createRequest.BrandNameEn),
		BrandStatus:   "DRAFT",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := brandService.brandRepo.Store(createBrand); err != nil {
		return err
	}
	return nil
}

func (brandService brandServiceImp) UpdateBrand(ctx *fiber.Ctx, config *config.Config, brandID int, updateRequest dto.UpdateBrandRequest) error {
	updateBrand, err := brandService.brandRepo.GetOneByID(brandID)
	if err != nil {
		return errors.New("brand not found")
	}
	file, _ := ctx.FormFile("brand_image_url")

	if file != nil {
		// old image delete
		if err := utils.DeleteFile(*updateBrand.BrandImageURL); err != nil {
			return err
		}
		// new image upload
		brandImageURL, err := utils.UploadFile(ctx, "brand_image_url", config.FolderConfig.PublicPath, "brand-images")
		if err != nil {
			return err
		}
		updateBrand.BrandImageURL = brandImageURL
	}

	updateBrand.BrandNameTk = updateRequest.BrandNameTk
	updateBrand.BrandNameRu = updateRequest.BrandNameRu
	updateBrand.BrandNameEn = updateRequest.BrandNameEn
	updateBrand.BrandStatus = updateRequest.BrandStatus
	updateBrand.BrandSlug = slug.Make(updateBrand.BrandNameEn)
	updateBrand.UpdatedAt = time.Now()

	if err := brandService.brandRepo.Update(updateBrand.ID, *updateBrand); err != nil {
		return err
	}
	return nil
}

func (brandService brandServiceImp) DeleteBrand(brandID int) error {
	if err := brandService.brandRepo.Destroy(brandID); err != nil {
		return err
	}
	return nil
}
