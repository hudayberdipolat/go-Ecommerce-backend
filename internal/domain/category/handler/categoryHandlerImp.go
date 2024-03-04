package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
	config          *config.Config
}

func NewCategoryHandler(service service.CategoryService, config *config.Config) CategoryHandler {
	return categoryHandlerImp{
		categoryService: service,
		config:          config,
	}
}

func (categoryHandler categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	category, err := categoryHandler.categoryService.GetOneCategoryWithID(categoryID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "category not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one category", category)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (categoryHandler categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categories, err := categoryHandler.categoryService.GetAllCategory()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "categories not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all categories", categories)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (categoryHandler categoryHandlerImp) Create(ctx *fiber.Ctx) error {

	var createCategoryRequest dto.CreateCategoryRequest
	// body parser
	if err := ctx.BodyParser(&createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// create category
	if err := categoryHandler.categoryService.CreateCategory(ctx, categoryHandler.config, createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create category error ", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusCreated, "category created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (categoryHandler categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	var updateCategoryRequest dto.UpdateCategoryRequest
	if err := ctx.BodyParser(&updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// update category
	err := categoryHandler.categoryService.UpdateCategory(ctx, categoryHandler.config, categoryID, updateCategoryRequest)
	if err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "update category error ", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusOK, "category updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (categoryHandler categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	if err := categoryHandler.categoryService.DeleteCategory(categoryID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "category delete error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "category deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
