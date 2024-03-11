package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	sliderRepo    repository.SliderRepository
	sliderService service.SliderService
	SliderHandler handler.SliderHandler
)

func SliderRequirmentCreator(db *gorm.DB, config *config.Config) {
	sliderRepo = repository.NewSliderRepository(db)
	sliderService = service.NewSliderService(sliderRepo)
	SliderHandler = handler.NewSliderHandler(sliderService, config)
}
