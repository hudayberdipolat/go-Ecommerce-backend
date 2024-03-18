package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type AppDataService interface {
	GetOne(appDataID int) (*models.AppData, error)
	CreateAppData(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateAppDataRequest) error
	UpdateAppData(ctx *fiber.Ctx, config *config.Config, appDataID int, updateRequest dto.UpdateAppDataRequest) error
}
