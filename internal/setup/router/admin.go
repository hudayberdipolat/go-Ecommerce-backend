package router

import (
	"github.com/gofiber/fiber/v2"
	aboutConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/constructor"
	BrendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	contactConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/constructor"
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
	productImageRoute.Delete("/:productImageID/delete", productImageConstructor.ProductImageHandler.Delete)

	// contact routes
	contactRoute := adminApi.Group("contact")
	contactRoute.Get("/:contactID", contactConstructor.ContactHandler.GetOne)
	contactRoute.Put("/:contactID/update", contactConstructor.ContactHandler.Update)

	// about routes

	aboutRoute := adminApi.Group("about")
	aboutRoute.Get("/:aboutID", aboutConstructor.AboutHandler.GetOne)
	aboutRoute.Get("/:aboutID/update", aboutConstructor.AboutHandler.Update)

}
