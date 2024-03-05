package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/handler"
	subCategoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	subCategoryRepo    subCategoryRepository.SubCategoryRepository
	subCategoryService service.SubCategoryService
	SubCategoryHandler handler.SubCategoryHandler
	// for category
	categoryRepo repository.CategoryRepository
)

func SubCategoryRequirmentCreator(db *gorm.DB, config *config.Config) {
	subCategoryRepo = subCategoryRepository.NewSubCategoryRespository(db)
	categoryRepo = repository.NewCategoryRepository(db)
	subCategoryService = service.NewSubCategoryService(subCategoryRepo, categoryRepo)
	SubCategoryHandler = handler.NewSubCategoryHandler(subCategoryService, config)
}
