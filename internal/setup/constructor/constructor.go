package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	aboutConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/constructor"
	adminConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/constructor"
	brendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	contactConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/constructor"
	permissionConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/constructor"
	productconstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	productImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/constructor"
	roleConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/constructor"
	userConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	adminConstructor.AdminRequirementsCreator(dependencies.DB)
	roleConstructor.RoleRequirementsCreator(dependencies.DB)
	permissionConstructor.PermissionRequirementsCreator(dependencies.DB)

	userConstructor.UserRequirementsCreator(dependencies.DB)

	categoryConstructor.CategoryRequirementsCreator(dependencies.DB)
	brendConstructor.BrendRequirementsCreator(dependencies.DB, dependencies.Config)
	productconstructor.ProductRequirementsCreator(dependencies.DB, dependencies.Config)
	productImageConstructor.ProductImageRequirementsCreator(dependencies.DB, dependencies.Config)
	contactConstructor.ContactRequirementsCreator(dependencies.DB)
	aboutConstructor.AboutRequirementsCreator(dependencies.DB)

}
