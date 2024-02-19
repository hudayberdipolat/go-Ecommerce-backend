package router

import (
	"github.com/gofiber/fiber/v2"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
)

func AdminRoutes(app *fiber.App) {
	// admin routes api
	adminApi := app.Group("/api/admin")

	// category routes
	categoryRoute := adminApi.Group("categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID", categoryConstructor.CategoryHandler.Delete)

}
