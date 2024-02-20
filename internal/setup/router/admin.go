package router

import (
	"github.com/gofiber/fiber/v2"
	BrendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	productImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/constructor"
)

func AdminRoutes(app *fiber.App) {
	// admin routes api
	adminApi := app.Group("/api/admin")

	// category routes
	categoryRoute := adminApi.Group("categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID/update", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID/delete", categoryConstructor.CategoryHandler.Delete)

	// brend routes
	brendRoute := adminApi.Group("brends")
	brendRoute.Get("/", BrendConstructor.BrendHandler.GetAll)
	brendRoute.Get("/:brendID", BrendConstructor.BrendHandler.GetOne)
	brendRoute.Post("/create", BrendConstructor.BrendHandler.Create)
	brendRoute.Put("/:brendID/update", BrendConstructor.BrendHandler.Update)
	brendRoute.Delete("/:brendID/delete", BrendConstructor.BrendHandler.Delete)

	// product routes
	productRoute := adminApi.Group("products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/:productID", productConstructor.ProductHandler.GetOne)
	productRoute.Post("/create", productConstructor.ProductHandler.Create)
	productRoute.Put("/:productID/update", productConstructor.ProductHandler.Update)
	productRoute.Delete("/:productID/delete", productConstructor.ProductHandler.Delete)

	// product images route
	productImageRoute := adminApi.Group("products/:productID/product-images")
	productImageRoute.Post("/create", productImageConstructor.ProductImageHandler.Create)
	productImageRoute.Delete("/:productImageID/delete", productImageConstructor.ProductImageHandler.Create)
}
