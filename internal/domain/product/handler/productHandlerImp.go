package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/service"
)

type productHandlerImp struct {
	productService service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return productHandlerImp{
		productService: service,
	}
}

func (productHandler productHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("product Handler Imp")
}

func (productHandler productHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("product Handler Imp")
}

func (productHandler productHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("product Handler Imp")
}

func (productHandler productHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("product Handler Imp")
}

func (productHandler productHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("product Handler Imp")
}
