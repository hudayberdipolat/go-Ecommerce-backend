package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type contactHandlerImp struct {
	contactService service.ContactService
	config         *config.Config
}

func NewContactHandler(service service.ContactService, config *config.Config) ContactHandler {
	return contactHandlerImp{
		contactService: service,
	}
}

func (contactHandler contactHandlerImp) GetContact(ctx *fiber.Ctx) error {
	panic("contact Handler IMP")
}

func (contactHandler contactHandlerImp) UpdateContact(ctx *fiber.Ctx) error {
	panic("contact Handler IMP")
}
