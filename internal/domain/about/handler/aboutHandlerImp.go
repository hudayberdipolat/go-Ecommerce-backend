package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/service"
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
	panic("about handler imp")
}

func (aboutHandler aboutHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("about handler imp")
}
