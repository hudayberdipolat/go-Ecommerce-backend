package service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageServiceImp struct {
	productImageRepo repository.ProductImageRepository
}

func NewProductImageService(repo repository.ProductImageRepository) ProductImageService {
	return productImageServiceImp{
		productImageRepo: repo,
	}
}

func (productImageSerive productImageServiceImp) CreateProductImage(ctx *fiber.Ctx, config *config.Config, request dto.ProductImageRequest) error {
	// get product
	getProduct, err := productImageSerive.productImageRepo.GetOneProduct(request.ProductID)
	if err != nil {
		return err
	}

	// image not empty validate

	image, _ := ctx.FormFile("product_image")

	if image == nil {
		return errors.New("product image empty ")
	}

	// file upload
	path, err := utils.UploadFile(ctx, "product_image", config.FolderConfig.PublicPath, "product-images")
	if err != nil {
		return err
	}

	request.ProductImage = *path

	createProductImage := models.ProductImage{
		ProductID:    getProduct.ID,
		ProductImage: request.ProductImage,
	}

	if err := productImageSerive.productImageRepo.Create(createProductImage); err != nil {
		return err
	}
	return nil

}

func (productImageSerive productImageServiceImp) Delete(productID, productImageID int) error {
	productImage, err := productImageSerive.productImageRepo.GetOneProductImage(productID, productImageID)
	if err != nil {
		return err
	}

	// image delete
	if err := utils.DeleteFile(productImage.ProductImage); err != nil {
		return err
	}

	if err := productImageSerive.productImageRepo.Delete(productImage.ProductID, productImage.ID); err != nil {
		return err
	}
	return nil

}
