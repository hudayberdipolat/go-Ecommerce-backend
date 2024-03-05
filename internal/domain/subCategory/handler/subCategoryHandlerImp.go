package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type subCategoryHandlerImp struct {
	subCategoryService service.SubCategoryService
	config             *config.Config
}

func NewSubCategoryHandler(service service.SubCategoryService, config *config.Config) SubCategoryHandler {
	return subCategoryHandlerImp{
		subCategoryService: service,
		config:             config,
	}
}

func (subCategoryHandler subCategoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("subCategory handler imp")
}

func (subCategoryHandler subCategoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("subCategory handler imp")
}

func (subCategoryHandler subCategoryHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("subCategory handler imp")
}

func (subCategoryHandler subCategoryHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("subCategory handler imp")
}

func (subCategoryHandler subCategoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("subCategory handler imp")
}
