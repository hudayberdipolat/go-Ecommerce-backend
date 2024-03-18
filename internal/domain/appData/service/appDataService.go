package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type AppDataService interface {
	GetOne() (*models.AppData, error)
	CreateAppData() error
	UpdateAppData(ctx *fiber.Ctx, config *config.Config, appDataID int, updateRequest dto.UpdateAppDataRequest) error
}
