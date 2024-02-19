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

func (p productHandlerImp) GetOne(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	product, err := p.productService.GetOneProduct(productID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "product not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one product", product)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p productHandlerImp) GetAll(ctx *fiber.Ctx) error {
	products, err := p.productService.GetAllProduct()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "products not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all product", products)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p productHandlerImp) Create(ctx *fiber.Ctx) error {
	var createProductRequest dto.CreateProductRequest

	// body parser
	if err := ctx.BodyParser(&createProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate request
	if err := validate.ValidateStruct(createProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create product

	if err := p.productService.CreateProduct(ctx, p.config, createProductRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create product error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "product created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p productHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateProductRequest dto.UpdateProductRequest
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	// body parser
	if err := ctx.BodyParser(&updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate request
	if err := validate.ValidateStruct(updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update product

	if err := p.productService.UpdateProduct(ctx, p.config, productID, updateProductRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "update product error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "product updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p productHandlerImp) Delete(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	if err := p.productService.DeleteProduct(productID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "deleted product error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "product deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
