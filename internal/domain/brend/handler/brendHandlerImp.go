package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/service"
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
	panic("brend handler imp")
}

func (b brendHandlerimp) GetAll(ctx *fiber.Ctx) error {
	panic("brend handler imp")
}

func (b brendHandlerimp) Create(ctx *fiber.Ctx) error {
	panic("brend handler imp")
}

func (b brendHandlerimp) Update(ctx *fiber.Ctx) error {
	panic("brend handler imp")
}

func (b brendHandlerimp) Delete(ctx *fiber.Ctx) error {
	panic("brend handler imp")
}
