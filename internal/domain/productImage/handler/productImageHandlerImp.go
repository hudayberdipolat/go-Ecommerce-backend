package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageHandlerImp struct {
	productImageService service.ProductImageService
	config              *config.Config
}

func NewProductImageHandler(service service.ProductImageService, config *config.Config) ProductImageHandler {
	return productImageHandlerImp{
		productImageService: service,
		config:              config,
	}
}

func (productImageHandler productImageHandlerImp) Create(ctx *fiber.Ctx) error {
	var productImageRequest dto.ProductImageRequest
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	if err := ctx.BodyParser(&productImageRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	productImageRequest.ProductID = productID

	err := productImageHandler.productImageService.CreateProductImage(ctx, productImageHandler.config, productImageRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "product image create error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "product Image created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (productImageHandler productImageHandlerImp) Delete(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	productImageID, _ := strconv.Atoi(ctx.Params("productImageID"))

	if err := productImageHandler.productImageService.Delete(productID, productImageID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "product image delete error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "product Image deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
