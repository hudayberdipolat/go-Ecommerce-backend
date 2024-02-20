package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	brendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	contactConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/constructor"
	productconstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	productImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirementsCreator(dependencies.DB)
	brendConstructor.BrendRequirementsCreator(dependencies.DB, dependencies.Config)
	productconstructor.ProductRequirementsCreator(dependencies.DB, dependencies.Config)
	productImageConstructor.ProductImageRequirementsCreator(dependencies.DB, dependencies.Config)
	contactConstructor.ContactRequirementsCreator(dependencies.DB)
}
