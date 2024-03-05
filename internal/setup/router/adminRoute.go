package router

import (
	"github.com/gofiber/fiber/v2"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	subCategoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/constructor"
)

func AdminRoutes(app *fiber.App) {

	adminApi := app.Group("/api/admin")

	// categories routes

	categoryRoute := adminApi.Group("categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID/update", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID/delete", categoryConstructor.CategoryHandler.Delete)

	// subCategory routes

	subCategoryRoute := adminApi.Group("/categories/:categoryID/subCategories")
	subCategoryRoute.Get("/", subCategoryConstructor.SubCategoryHandler.GetAll)
	subCategoryRoute.Get("/:subCategoryID", subCategoryConstructor.SubCategoryHandler.GetOne)
	subCategoryRoute.Post("/create", subCategoryConstructor.SubCategoryHandler.Create)
	subCategoryRoute.Put("/:subCategoryID/update", subCategoryConstructor.SubCategoryHandler.Update)
	subCategoryRoute.Delete("/:subCategoryID/delete", subCategoryConstructor.SubCategoryHandler.Delete)

}
