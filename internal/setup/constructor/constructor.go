package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirmentCreator(dependencies.DB, dependencies.Config)
}
