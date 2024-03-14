package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type sliderHandlerImp struct {
	sliderService service.SliderService
	config        *config.Config
}

func NewSliderHandler(service service.SliderService, config *config.Config) SliderHandler {
	return sliderHandlerImp{
		sliderService: service,
		config:        config,
	}
}

func (sliderHandler sliderHandlerImp) GetAll(ctx *fiber.Ctx) error {
	sliders, err := sliderHandler.sliderService.GetAllSlider()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "sliders not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all sliders", sliders)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (sliderHandler sliderHandlerImp) GetOne(ctx *fiber.Ctx) error {
	sliderID, _ := strconv.Atoi(ctx.Params("sliderID"))
	slider, err := sliderHandler.sliderService.GetOneSlider(sliderID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "slider not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one slider", slider)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (sliderHandler sliderHandlerImp) Create(ctx *fiber.Ctx) error {
	var createSliderRequest dto.CreateSliderRequest
	// body parser
	if err := ctx.BodyParser(&createSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(createSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	if err := sliderHandler.sliderService.CreateSlider(ctx, sliderHandler.config, createSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "slider can't created", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "slider created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (sliderHandler sliderHandlerImp) UpdateStatus(ctx *fiber.Ctx) error {
	var updateSliderRequest dto.UpdateSliderStatus
	sliderID, _ := strconv.Atoi(ctx.Params("sliderID"))
	// body parser
	if err := ctx.BodyParser(&updateSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(updateSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	if err := sliderHandler.sliderService.UpdateSliderStatus(sliderID, updateSliderRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "slider can't created", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "slider updated successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (sliderHandler sliderHandlerImp) Delete(ctx *fiber.Ctx) error {
	sliderID, _ := strconv.Atoi(ctx.Params("sliderID"))
	if err := sliderHandler.sliderService.DeleteSlider(sliderID); err != nil {
		errResponse := response.Error(http.StatusNotFound, "slider can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "slider deleted successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}
