package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/service"
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
	panic("contact handler imp")
}

func (contactHandler contactHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("contact handler imp")
}
