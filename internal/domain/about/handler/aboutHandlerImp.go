package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type aboutHandlerImp struct {
	aboutService service.AboutService
}

func NewAboutHandler(service service.AboutService) AboutHandler {
	return aboutHandlerImp{
		aboutService: service,
	}
}

func (aboutHandler aboutHandlerImp) GetOne(ctx *fiber.Ctx) error {
	aboutID, _ := strconv.Atoi(ctx.Params("aboutID"))

	about, err := aboutHandler.aboutService.GetOneAbout(aboutID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "about not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, " get about data", about)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (aboutHandler aboutHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateAboutRequest dto.AboutRequest
	aboutID, _ := strconv.Atoi(ctx.Params("aboutID"))
	// body parser
	if err := ctx.BodyParser(&updateAboutRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(updateAboutRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// update about data

	if err := aboutHandler.aboutService.UpdateAbout(aboutID, updateAboutRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "about update error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "about updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
