package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type brandHandlerImp struct {
	brandService service.BrandService
	config       *config.Config
}

func NewBrandHandler(brandService service.BrandService, config *config.Config) BrandHandler {
	return brandHandlerImp{
		brandService: brandService,
		config:       config,
	}
}

func (brandHandler brandHandlerImp) GetOne(ctx *fiber.Ctx) error {
	brandID, _ := strconv.Atoi(ctx.Params("brandID"))
	brand, err := brandHandler.brandService.GetOneBrandByID(brandID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "brand not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get One brand", brand)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (brandHandler brandHandlerImp) GetAll(ctx *fiber.Ctx) error {
	brands, err := brandHandler.brandService.GetAllBrand()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "brands not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all brands", brands)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (brandHandler brandHandlerImp) Create(ctx *fiber.Ctx) error {
	var createBrandRequest dto.CreateBrandRequest
	// body parser
	if err := ctx.BodyParser(&createBrandRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data

	if err := validate.ValidateStruct(createBrandRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create brand

	if err := brandHandler.brandService.CreateBrand(ctx, brandHandler.config, createBrandRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "brand not created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "brand created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (brandHandler brandHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateBrandRequest dto.UpdateBrandRequest
	brandID, _ := strconv.Atoi(ctx.Params("brandID"))

	// body parser

	if err := ctx.BodyParser(&updateBrandRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data

	if err := validate.ValidateStruct(updateBrandRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update brand

	if err := brandHandler.brandService.UpdateBrand(ctx, brandHandler.config, brandID, updateBrandRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "brand not updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// return response
	successResponse := response.Success(http.StatusOK, "brand updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (brandHandler brandHandlerImp) Delete(ctx *fiber.Ctx) error {
	brandID, _ := strconv.Atoi(ctx.Params("brandID"))

	if err := brandHandler.brandService.DeleteBrand(brandID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "brand not deleted", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// return response
	successResponse := response.Success(http.StatusOK, "brand deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
