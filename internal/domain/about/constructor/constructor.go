package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/service"
	"gorm.io/gorm"
)

var (
	aboutRepo    repository.AboutRepository
	aboutService service.AboutService
	AboutHandler handler.AboutHandler
)

func AboutRequirementsCreator(db *gorm.DB) {
	aboutRepo = repository.NewAboutRepository(db)
	aboutService = service.NewAboutService(aboutRepo)
	AboutHandler = handler.NewAboutHandler(aboutService)
}
