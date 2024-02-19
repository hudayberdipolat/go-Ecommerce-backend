package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/service"
	"gorm.io/gorm"
)

var (
	categoryRepo    repository.CategoryRepository
	categoryServive service.CategoryService
	CategoryHandler handler.CategoryHandler
)

func CategoryRequirementsCreator(db *gorm.DB) {
	categoryRepo = repository.NewCategoryRepository(db)
	categoryServive = service.NewCategoryService(categoryRepo)
	CategoryHandler = handler.NewCategoryHandler(categoryServive)
}
