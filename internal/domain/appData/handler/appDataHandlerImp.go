package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type appDataHandlerImp struct {
	appDataService service.AppDataService
	config         *config.Config
}

func NewAppDataHandler(service service.AppDataService, config *config.Config) AppDataHandler {
	return appDataHandlerImp{
		appDataService: service,
		config:         config,
	}
}

func (appDataHandler appDataHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("App Data Handler IMP")
}

func (appDataHandler appDataHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("App Data Handler IMP")
}

func (appDataHandler appDataHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("App Data Handler IMP")
}

func (appDataHandler appDataHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("App Data Handler IMP")
}
