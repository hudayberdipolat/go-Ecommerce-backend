package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	categoryRepo    repository.CategoryRepository
	categoryService service.CategoryService
	CategoryHandler handler.CategoryHandler
)

func CategoryRequirmentCreator(db *gorm.DB, config *config.Config) {
	categoryRepo = repository.NewCategoryRepository(db)
	categoryService = service.NewCategoryService(categoryRepo)
	CategoryHandler = handler.NewCategoryHandler(categoryService, config)
}
