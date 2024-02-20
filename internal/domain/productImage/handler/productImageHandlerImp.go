package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/service"
)

type productImageHandlerImp struct {
	productImageService service.ProductImageService
}

func NewProductImageHandler(service service.ProductImageService) ProductImageHandler {
	return productImageHandlerImp{
		productImageService: service,
	}
}

func (productImageHandler productImageHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("product image handler imp")
}

func (productImageHandler productImageHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("product image handler imp")
}
