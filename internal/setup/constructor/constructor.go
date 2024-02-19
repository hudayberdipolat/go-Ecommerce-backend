package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	brendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirementsCreator(dependencies.DB)
	brendConstructor.BrendRequirementsCreator(dependencies.DB)
}
