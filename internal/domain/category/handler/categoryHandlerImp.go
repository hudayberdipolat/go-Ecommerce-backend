package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return categoryHandlerImp{
		categoryService: service,
	}
}

func (handler categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	category, err := handler.categoryService.FindOneCategory(categoryID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "category not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one category", category)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (handler categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categories, err := handler.categoryService.FindAllCategory()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "categories not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all categories", categories)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (handler categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	var createCategoryRequest dto.CreateCategoryRequest
	// body parser request
	if err := ctx.BodyParser(&createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate request
	if err := validate.ValidateStruct(createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create category
	if err := handler.categoryService.CreateCategory(createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create category error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	// return response
	successResponse := response.Success(http.StatusCreated, "category created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (handler categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateCategoryRequest dto.UpdateCategoryRequest
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	// body parser request
	if err := ctx.BodyParser(&updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate request
	if err := validate.ValidateStruct(updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update category

	if err := handler.categoryService.UpdateCategory(categoryID, updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "update category error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	// return success response
	successResponse := response.Success(http.StatusOK, "category updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (handler categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))

	if err := handler.categoryService.DeleteCategory(categoryID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "delete category error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "category deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
