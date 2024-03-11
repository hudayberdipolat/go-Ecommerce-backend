package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/service"
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
	panic("sliderHandler IMP")
}

func (sliderHandler sliderHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("sliderHandler IMP")
}

func (sliderHandler sliderHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("sliderHandler IMP")
}

func (sliderHandler sliderHandlerImp) UpdateStatus(ctx *fiber.Ctx) error {
	panic("sliderHandler IMP")
}

func (sliderHandler sliderHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("sliderHandler IMP")
}
