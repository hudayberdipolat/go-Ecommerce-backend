package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type subCategoryHandlerImp struct {
	subCategoryService service.SubCategoryService
	config             *config.Config
}

func NewSubCategoryHandler(service service.SubCategoryService, config *config.Config) SubCategoryHandler {
	return subCategoryHandlerImp{
		subCategoryService: service,
		config:             config,
	}
}

func (subCategoryHandler subCategoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	subCategoryID, _ := strconv.Atoi(ctx.Params("subCategoryID"))

	subCategory, err := subCategoryHandler.subCategoryService.GetOneSubCategory(categoryID, subCategoryID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "subCategory not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one subCategory", subCategory)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (subCategoryHandler subCategoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	subCategories, err := subCategoryHandler.subCategoryService.GetAllSubCategory(categoryID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "subCategories not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all subCategory", subCategories)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (subCategoryHandler subCategoryHandlerImp) Create(ctx *fiber.Ctx) error {
	var createSubCategoryRequest dto.CreateSubCategoryRequest
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))

	// body parser
	if err := ctx.BodyParser(&createSubCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(createSubCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create subCategory
	err := subCategoryHandler.subCategoryService.CreateSubCategory(ctx, subCategoryHandler.config, categoryID, createSubCategoryRequest)
	if err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create subCategory Error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "created successfully subCategory", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (subCategoryHandler subCategoryHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateSubCategoryRequest dto.UpdateSubCategoryRequest
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	subCategoryID, _ := strconv.Atoi(ctx.Params("subCategoryID"))
	// body parser
	if err := ctx.BodyParser(&updateSubCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateSubCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update subCategory

	err := subCategoryHandler.subCategoryService.UpdateSubCategory(ctx, subCategoryHandler.config, categoryID, subCategoryID, updateSubCategoryRequest)
	if err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "update subCategory Error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "updated successfully subCategory", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (subCategoryHandler subCategoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	subCategoryID, _ := strconv.Atoi(ctx.Params("subCategoryID"))

	if err := subCategoryHandler.subCategoryService.DeleteSubCategory(categoryID, subCategoryID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "delete subCategory Error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "deleted successfully subCategory", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
