package constructor

import (
	brandRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
	categoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/handler"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/service"
	subCategoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	productRepo     productRepository.ProductRepository
	categoryRepo    categoryRepository.CategoryRepository
	subCategoryRepo subCategoryRepository.SubCategoryRepository
	brandRepo       brandRepository.BrandRepository

	productService service.ProductService
	ProductHandler handler.ProductHandler
)

func ProductRequirmentCreator(db *gorm.DB, config *config.Config) {

	productRepo = productRepository.NewProductRepository(db)
	categoryRepo = categoryRepository.NewCategoryRepository(db)
	subCategoryRepo = subCategoryRepository.NewSubCategoryRespository(db)
	brandRepo = brandRepository.NewBrandRepository(db)
	productService = service.NewProductService(productRepo, categoryRepo, subCategoryRepo, brandRepo)
	ProductHandler = handler.NewProductHandler(productService, config)

}
