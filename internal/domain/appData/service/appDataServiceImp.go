package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
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

func (appDataService appDataServiceImp) GetOne(appDataID int) (*models.AppData, error) {
	appData, err := appDataService.appDataRepo.FindOne(appDataID)
	if err != nil {
		return nil, err
	}
	return appData, nil
}

func (appDataService appDataServiceImp) CreateAppData(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateAppDataRequest) error {

	appLogo, err := utils.UploadFile(ctx, "app_logo", config.FolderConfig.PublicPath, "app-data-images")
	if err != nil {
		return err
	}
	playMarket, err := utils.UploadFile(ctx, "play_market", config.FolderConfig.PublicPath, "app-data-images")
	if err != nil {
		return err
	}
	appStore, err := utils.UploadFile(ctx, "app_store", config.FolderConfig.PublicPath, "app-data-images")
	if err != nil {
		return err
	}
	qrCode, err := utils.UploadFile(ctx, "qr_code", config.FolderConfig.PublicPath, "app-data-images")
	if err != nil {
		return err
	}

	appData := models.AppData{
		AppLogo:    appLogo,
		PlayMarket: playMarket,
		AppStore:   appStore,
		QrCode:     qrCode,
	}

	if err := appDataService.appDataRepo.Create(appData); err != nil {
		if err := utils.DeleteFile(*appLogo); err != nil {
			return err
		}
		if err := utils.DeleteFile(*playMarket); err != nil {
			return err
		}
		if err := utils.DeleteFile(*appStore); err != nil {
			return err
		}
		if err := utils.DeleteFile(*qrCode); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (appDataService appDataServiceImp) UpdateAppData(ctx *fiber.Ctx, config *config.Config, appDataID int, updateRequest dto.UpdateAppDataRequest) error {
	appData, err := appDataService.appDataRepo.FindOne(appDataID)
	if err != nil {
		return err
	}
	appLogoImage, _ := ctx.FormFile("app_logo")
	if appLogoImage != nil {
		// old logo delete new image upload
		if err := utils.DeleteFile(*appData.AppLogo); err != nil {
			return err
		}
		// new image upload
		appLogo, err := utils.UploadFile(ctx, "app_logo", config.FolderConfig.PublicPath, "app-data-images")
		if err != nil {
			return err
		}
		appData.AppLogo = appLogo
	}
	playMarketImage, _ := ctx.FormFile("play_market")
	if playMarketImage != nil {
		// old logo delete new image upload
		if err := utils.DeleteFile(*appData.PlayMarket); err != nil {
			return err
		}
		// new image upload
		playMarket, err := utils.UploadFile(ctx, "play_market", config.FolderConfig.PublicPath, "app-data-images")
		if err != nil {
			return err
		}
		appData.PlayMarket = playMarket
	}
	appStoreImage, _ := ctx.FormFile("app_store")

	if appStoreImage != nil {
		// old logo delete new image upload
		if err := utils.DeleteFile(*appData.AppStore); err != nil {
			return err
		}
		// new image upload
		appStore, err := utils.UploadFile(ctx, "app_store", config.FolderConfig.PublicPath, "app-data-images")
		if err != nil {
			return err
		}
		appData.AppStore = appStore
	}

	qrCodeImage, _ := ctx.FormFile("qr_code")
	if qrCodeImage != nil {
		// old logo delete new image upload
		if err := utils.DeleteFile(*appData.QrCode); err != nil {
			return err
		}
		// new image upload
		qrCode, err := utils.UploadFile(ctx, "qr_code", config.FolderConfig.PublicPath, "app-data-images")
		if err != nil {
			return err
		}
		appData.AppLogo = qrCode
	}
	if err := appDataService.appDataRepo.Update(appData.ID, *appData); err != nil {
		return err
	}
	return nil
}
