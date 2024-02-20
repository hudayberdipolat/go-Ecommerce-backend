package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	productImageRepo    repository.ProductImageRepository
	productImageService service.ProductImageService
	ProductImageHandler handler.ProductImageHandler
)

func ProductImageRequirementsCreator(db *gorm.DB, config *config.Config) {
	productImageRepo = repository.NewProductImageRepository(db)
	productImageService = service.NewProductImageService(productImageRepo)
	ProductImageHandler = handler.NewProductImageHandler(productImageService)
}
