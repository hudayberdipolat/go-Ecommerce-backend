package router

import (
	"github.com/gofiber/fiber/v2"
	brendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	userConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/constructor"
)

func UserRoutes(app *fiber.App) {

	// user api routes
	userApi := app.Group("api/front")

	authRoute := userApi.Group("auth")
	authRoute.Post("/register", userConstructor.UserHandler.Register)
	authRoute.Post("/login", userConstructor.UserHandler.Login)

	// categories
	categoryRoute := userApi.Group("categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	// category one get with  category slug
	// categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetOne)

	// products
	productRoute := userApi.Group("products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/:productSlug", productConstructor.ProductHandler.GetOneProduct)

	// brends

	brendRoute := userApi.Group("brends")
	brendRoute.Get("/", brendConstructor.BrendHandler.GetAll)
	// brend get one with slug
	// brendRoute.Get("/:brendSlug", brendConstructor.BrendHandler.GetAll)
}
