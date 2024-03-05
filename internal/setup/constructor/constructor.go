package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	brandConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	subCategoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirmentCreator(dependencies.DB, dependencies.Config)
	subCategoryConstructor.SubCategoryRequirmentCreator(dependencies.DB, dependencies.Config)
	brandConstructor.BrandRequirmentCreator(dependencies.DB, dependencies.Config)
}
