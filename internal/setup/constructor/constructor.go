package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	adminConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/constructor"
	brandConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	pImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	sliderConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/constructor"
	subCategoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/constructor"
	userConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirmentCreator(dependencies.DB, dependencies.Config)
	subCategoryConstructor.SubCategoryRequirmentCreator(dependencies.DB, dependencies.Config)
	brandConstructor.BrandRequirmentCreator(dependencies.DB, dependencies.Config)
	productConstructor.ProductRequirmentCreator(dependencies.DB, dependencies.Config)
	pImageConstructor.ProductImageRequirmentCreator(dependencies.DB, dependencies.Config)
	userConstructor.UserRequirmentCreator(dependencies.DB)
	adminConstructor.AdminRequirmentCreator(dependencies.DB, dependencies.Config)
	sliderConstructor.SliderRequirmentCreator(dependencies.DB, dependencies.Config)
}
