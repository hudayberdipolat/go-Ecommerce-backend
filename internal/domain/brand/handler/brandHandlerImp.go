package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type brandHandlerImp struct {
	brandService service.BrandService
	config       *config.Config
}

func NewBrandHandler(brandService service.BrandService, config *config.Config) BrandHandler {
	return brandHandlerImp{
		brandService: brandService,
		config:       config,
	}
}

func (brandHandler brandHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("brand handler Imp")
}

func (brandHandler brandHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("brand handler Imp")
}

func (brandHandler brandHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("brand handler Imp")
}

func (brandHandler brandHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("brand handler Imp")
}

func (brandHandler brandHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("brand handler Imp")
}
