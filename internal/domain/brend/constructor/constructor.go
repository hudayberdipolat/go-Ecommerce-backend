package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/service"
	"gorm.io/gorm"
)

var (
	brendRepo    repository.BrendRepository
	brendService service.BrendService
	BrendHandler handler.BrendHandler
)

func BrendRequirementsCreator(db *gorm.DB) {
	brendRepo = repository.NewBrendRepository(db)
	brendService = service.NewBrendService(brendRepo)
	BrendHandler = handler.NewBrendHandler(brendService)
}
