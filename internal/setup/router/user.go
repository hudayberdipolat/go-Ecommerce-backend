package router

import (
	"github.com/gofiber/fiber/v2"
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

	// products
	productRoute := app.Group("products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/", productConstructor.ProductHandler.GetOneProduct)

	// brends

}
