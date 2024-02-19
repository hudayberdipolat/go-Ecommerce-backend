package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	productRepo    repository.ProductRepository
	productService service.ProductService
	ProductHandler handler.ProductHandler
)

func ProductRequirementsCreator(db *gorm.DB, config *config.Config) {
	productRepo = repository.NewProductRepository(db)
	productService = service.NewProductService(productRepo)
	ProductHandler = handler.NewProductHandler(productService, config)
}
