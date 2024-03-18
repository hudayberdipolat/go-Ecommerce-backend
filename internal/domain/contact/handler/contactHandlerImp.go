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

func (contactHandler contactHandlerImp) GetContact(ctx *fiber.Ctx) error {
	contactID, _ := strconv.Atoi(ctx.Params("contactID"))
	contact, err := contactHandler.contactService.GetContact(contactID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "contact not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get contact data", contact)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (contactHandler contactHandlerImp) UpdateContact(ctx *fiber.Ctx) error {
	contactID, _ := strconv.Atoi(ctx.Params("contactID"))
	var updateContactRequest dto.UpdateContactRequest

	// body parser
	if err := ctx.BodyParser(&updateContactRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateContactRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update contact
	if err := contactHandler.contactService.UpdateContact(contactID, updateContactRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "contact can't updated", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "contact updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
