package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/handler"
	pImageRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/service"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	pImageRepo          pImageRepository.ProductImageRepository
	productRepo         productRepository.ProductRepository
	pImageService       service.ProductImageService
	ProductImageHandler handler.ProductImageHandler
)

func ProductImageRequirmentCreator(db *gorm.DB, config *config.Config) {
	pImageRepo = pImageRepository.NewProductImageRepository(db)
	productRepo = productRepository.NewProductRepository(db)
	pImageService = service.NewProductImageService(pImageRepo, productRepo)
	ProductImageHandler = handler.NewProductImageHandler(pImageService, config)
}
