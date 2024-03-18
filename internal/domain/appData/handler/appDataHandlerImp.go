package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type appDataHandlerImp struct {
	appDataService service.AppDataService
	config         *config.Config
}

func NewAppDataHandler(service service.AppDataService, config *config.Config) AppDataHandler {
	return appDataHandlerImp{
		appDataService: service,
		config:         config,
	}
}

func (appDataHandler appDataHandlerImp) GetOne(ctx *fiber.Ctx) error {
	appDataID, _ := strconv.Atoi(ctx.Params("appDataID"))

	appData, err := appDataHandler.appDataService.GetOne(appDataID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "appData not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get app Data", appData)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (appDataHandler appDataHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateAppDataRequest

	// body parser
	if err := ctx.BodyParser(&createRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(createRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// create app data

	if err := appDataHandler.appDataService.CreateAppData(ctx, appDataHandler.config, createRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create appData error ", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "appData created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (appDataHandler appDataHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateAppDataRequest
	appDataID, _ := strconv.Atoi(ctx.Params("appDataID"))

	// body parser
	if err := ctx.BodyParser(&updateRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate error
	if err := validate.ValidateStruct(updateRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// update
	err := appDataHandler.appDataService.UpdateAppData(ctx, appDataHandler.config, appDataID, updateRequest)
	if err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "can't updated appData ", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusOK, "appData updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
