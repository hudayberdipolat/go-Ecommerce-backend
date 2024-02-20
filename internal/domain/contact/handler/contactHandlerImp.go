package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type contactHandlerImp struct {
	contactService service.ContactService
}

func NewContactHandler(service service.ContactService) ContactHandler {
	return contactHandlerImp{
		contactService: service,
	}
}

func (contactHandler contactHandlerImp) GetOne(ctx *fiber.Ctx) error {
	contactID, _ := strconv.Atoi(ctx.Params("contactID"))

	contact, err := contactHandler.contactService.GetOneContact(contactID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "contact get error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusOK, "contact data", contact)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (contactHandler contactHandlerImp) Update(ctx *fiber.Ctx) error {
	contactID, _ := strconv.Atoi(ctx.Params("contactID"))

	var contactRequest dto.ContactRequest

	if err := ctx.BodyParser(&contactRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate

	if err := validate.ValidateStruct(&contactRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "contact validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	phoneNumberValidate := validate.PhoneNumberValidate(contactRequest.PhoneNumber)
	if !phoneNumberValidate {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi!!!", "Nädogry telefon belgi!!!", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate phone number

	if err := contactHandler.contactService.UpdateContact(contactID, contactRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "contact update error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "contact updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}
