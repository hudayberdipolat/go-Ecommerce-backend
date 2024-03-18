package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type appDataServiceImp struct {
	appDataRepo repository.AppDataRepository
}

func NewAppDataService(repo repository.AppDataRepository) AppDataService {
	return appDataServiceImp{
		appDataRepo: repo,
	}
}

func (appDataService appDataServiceImp) GetOne() (*models.AppData, error) {
	panic("app data Service IMP")
}

func (appDataService appDataServiceImp) CreateAppData() error {
	panic("app data Service IMP")
}

func (appDataService appDataServiceImp) UpdateAppData(ctx *fiber.Ctx, config *config.Config, appDataID int, updateRequest dto.UpdateAppDataRequest) error {
	panic("app data Service IMP")
}
