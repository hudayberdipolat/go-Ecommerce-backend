package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageHandlerImp struct {
	pImageService service.ProductImageService
	config        *config.Config
}

func NewProductImageHandler(service service.ProductImageService, conf *config.Config) ProductImageHandler {
	return productImageHandlerImp{
		pImageService: service,
		config:        conf,
	}
}

func (pImageHandler productImageHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("productImage Handler imp")
}

func (pImageHandler productImageHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("productImage Handler imp")
}

func (pImageHandler productImageHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("productImage Handler imp")
}

func (pImageHandler productImageHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("productImage Handler imp")
}
