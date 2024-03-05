package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	brandRepo    repository.BrandRepository
	brandService service.BrandService
	BrandHandler handler.BrandHandler
)

func BrandRequirmentCreator(db *gorm.DB, config *config.Config) {
	brandRepo = repository.NewBrandRepository(db)
	brandService = service.NewBrandService(brandRepo)
	BrandHandler = handler.NewBrandHandler(brandService, config)
}
