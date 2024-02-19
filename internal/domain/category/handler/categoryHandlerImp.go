package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/service"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return categoryHandlerImp{
		categoryService: service,
	}
}

func (handler categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (handler categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (handler categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (handler categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (handler categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("category handler imp")
}
