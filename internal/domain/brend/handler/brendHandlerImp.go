package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type brendHandlerimp struct {
	brendService service.BrendService
	config       *config.Config
}

func NewBrendHandler(service service.BrendService, config *config.Config) BrendHandler {
	return brendHandlerimp{
		brendService: service,
		config:       config,
	}
}

func (b brendHandlerimp) GetOne(ctx *fiber.Ctx) error {
	brendID, _ := strconv.Atoi(ctx.Params("brendID"))
	brend, err := b.brendService.GetOneBrend(brendID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "brend not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one brend", brend)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (b brendHandlerimp) GetAll(ctx *fiber.Ctx) error {
	brends, err := b.brendService.GetAllBrend()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "brends not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all brend", brends)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (b brendHandlerimp) Create(ctx *fiber.Ctx) error {
	var createBrendRequest dto.CreateBrendRequest

	// body parser
	if err := ctx.BodyParser(&createBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate

	if err := validate.ValidateStruct(createBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// create brend

	if err := b.brendService.CreateBrend(ctx, *b.config, createBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "create brend error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusOK, "Brend created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (b brendHandlerimp) Update(ctx *fiber.Ctx) error {
	var updateBrendRequest dto.UpdateBrendRequest
	brendID, _ := strconv.Atoi(ctx.Params("brendID"))
	// body parser
	if err := ctx.BodyParser(&updateBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate

	if err := validate.ValidateStruct(updateBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// update brend
	if err := b.brendService.UpdateBrend(ctx, *b.config, brendID, updateBrendRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "update brend error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusOK, "Brend updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (b brendHandlerimp) Delete(ctx *fiber.Ctx) error {
	brendID, _ := strconv.Atoi(ctx.Params("brendID"))

	if err := b.brendService.DeleteBrend(brendID); err != nil {
		errResponse := response.Error(http.StatusNotFound, "delete brend error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "Brend deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (b brendHandlerimp) GetOneBrend(ctx *fiber.Ctx) error {
	brendSlug := ctx.Params("brendSlug")
	brend, err := b.brendService.GetOneBrendWithSlug(brendSlug)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "brend not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one brend", brend)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
