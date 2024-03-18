package service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/dto"
	pImageRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/repository"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageServiceImp struct {
	pImageRepo  pImageRepository.ProductImageRepository
	productRepo productRepository.ProductRepository
}

func NewProductImageService(repo pImageRepository.ProductImageRepository, pRepo productRepository.ProductRepository) ProductImageService {
	return productImageServiceImp{
		pImageRepo:  repo,
		productRepo: pRepo,
	}
}

func (pImageService productImageServiceImp) GetAllProductImages(productID int) ([]dto.ProductImageResponse, error) {

	product, err := pImageService.productRepo.GetOneByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	productImages, err := pImageService.pImageRepo.FindAll(product.ID)
	if err != nil {
		return nil, err
	}

	productImageResponses := dto.NewGetAllProductImages(productImages)
	return productImageResponses, nil
}

func (pImageService productImageServiceImp) GetOneProductImage(productID, productImageID int) (*dto.ProductImageResponse, error) {
	product, err := pImageService.productRepo.GetOneByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	productImage, err := pImageService.pImageRepo.FindOne(product.ID, productImageID)
	if err != nil {
		return nil, err
	}

	productImageResponse := dto.NewGetOneProductImage(productImage)
	return &productImageResponse, nil

}

func (pImageService productImageServiceImp) CreateProductImage(ctx *fiber.Ctx, config *config.Config, productID int, request dto.CreateProductImageRequest) error {
	product, err := pImageService.productRepo.GetOneByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	// file upload
	imageURL, err := utils.UploadFile(ctx, "image_url", config.FolderConfig.PublicPath, "product-images")
	if err != nil {
		return err
	}

	// create product Image
	productImage := models.ProductImage{
		ImageURL:  imageURL,
		ProductID: product.ID,
	}

	if err := pImageService.pImageRepo.Store(productImage); err != nil {
		if err := utils.DeleteFile(*imageURL); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (pImageService productImageServiceImp) DeleteProductImage(productID, productImageID int) error {
	product, err := pImageService.productRepo.GetOneByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	productImage, err := pImageService.pImageRepo.FindOne(product.ID, productImageID)
	if err != nil {
		return err
	}

	if deleteImage := utils.DeleteFile(*productImage.ImageURL); deleteImage != nil {
		return deleteImage
	}

	if err := pImageService.pImageRepo.Destroy(product.ID, productImage.ID); err != nil {
		return err
	}
	return nil
}
