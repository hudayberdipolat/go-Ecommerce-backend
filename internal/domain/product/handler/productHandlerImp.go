package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productHandlerImp struct {
	productService service.ProductService
	config         *config.Config
}

func NewProductHandler(service service.ProductService, config *config.Config) ProductHandler {
	return productHandlerImp{
		productService: service,
		config:         config,
	}
}

func (productHandler productHandlerImp) GetOne(ctx *fiber.Ctx) error {

	productID, _ := strconv.Atoi(ctx.Params("productID"))

	product, err := productHandler.productService.GetOneProductByID(productID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "product not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one product", product)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (productHandler productHandlerImp) GetAll(ctx *fiber.Ctx) error {

	products, err := productHandler.productService.GetAllProduct()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "products not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all product", products)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (productHandler productHandlerImp) Create(ctx *fiber.Ctx) error {
	var createProductRequest dto.CreateProductRequest

	// body parser
	if err := ctx.BodyParser(&createProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(createProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create product
	if err := productHandler.productService.CreateProduct(ctx, productHandler.config, createProductRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "product not created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "product created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (productHandler productHandlerImp) Update(ctx *fiber.Ctx) error {

	var updateProductRequest dto.UpdateProductRequest
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	// body parser
	if err := ctx.BodyParser(&updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := productHandler.productService.UpdateProduct(ctx, productHandler.config, productID, updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "product not updated", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "product updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (productHandler productHandlerImp) Delete(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	if err := productHandler.productService.DeleteProduct(productID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "product not deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "product deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// FOR FRONT

func (productHandler productHandlerImp) GetOneProduct(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("productSlug")
	product, err := productHandler.productService.GetOneProductBySlug(productSlug)
	if err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "product not found", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get product", product)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
