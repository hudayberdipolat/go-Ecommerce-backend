package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/appData/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	appDataRepo    repository.AppDataRepository
	appDataService service.AppDataService
	AppDataHandler handler.AppDataHandler
)

func AppDataRequirmentCreator(db *gorm.DB, config *config.Config) {
	appDataRepo = repository.NewAppDataRepository(db)
	appDataService = service.NewAppDataService(appDataRepo)
	AppDataHandler = handler.NewAppDataHandler(appDataService, config)
}
