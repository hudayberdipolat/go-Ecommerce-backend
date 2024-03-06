package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageHandlerImp struct {
	pImageService service.ProductImageService
	config        *config.Config
}

func NewProductImageHandler(service service.ProductImageService, conf *config.Config) ProductImageHandler {
	return productImageHandlerImp{
		pImageService: service,
		config:        conf,
	}
}

func (pImageHandler productImageHandlerImp) GetOne(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	productImageID, _ := strconv.Atoi(ctx.Params("productImageID"))

	productImage, err := pImageHandler.pImageService.GetOneProductImage(productID, productImageID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "product image not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one product image", productImage)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (pImageHandler productImageHandlerImp) GetAll(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	productImages, err := pImageHandler.pImageService.GetAllProductImages(productID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "product images not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all product images", productImages)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (pImageHandler productImageHandlerImp) Create(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	var createProductImage dto.CreateProductImageRequest

	// body parser
	if err := ctx.BodyParser(&createProductImage); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(createProductImage); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create image

	if err := pImageHandler.pImageService.CreateProductImage(ctx, pImageHandler.config, productID, createProductImage); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "not created product images", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "product-image created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (pImageHandler productImageHandlerImp) Delete(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	productImageID, _ := strconv.Atoi(ctx.Params("productImageID"))

	if err := pImageHandler.pImageService.DeleteProductImage(productID, productImageID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "not created product images", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "product-image deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
