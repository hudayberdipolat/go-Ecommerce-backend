package router

import (
	"github.com/gofiber/fiber/v2"
	adminConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/constructor"
	brandConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	pImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	sliderConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/constructor"
	subCategoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/constructor"
)

func AdminRoutes(app *fiber.App) {

	adminApi := app.Group("/api/admin")

	// admin Auth routes

	// admin CRUD routes

	adminRoute := adminApi.Group("admins")
	adminRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminRoute.Post("/create", adminConstructor.AdminHandler.Create)
	adminRoute.Put("/:adminID/update", adminConstructor.AdminHandler.UpdateData)
	adminRoute.Put("/:adminID/update-password", adminConstructor.AdminHandler.UpdatePassword)
	adminRoute.Delete("/:adminID/delete", adminConstructor.AdminHandler.Delete)
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

	// brand routes

	brandRoute := adminApi.Group("/brands")
	brandRoute.Get("/", brandConstructor.BrandHandler.GetAll)
	brandRoute.Get("/:brandID", brandConstructor.BrandHandler.GetOne)
	brandRoute.Post("/create", brandConstructor.BrandHandler.Create)
	brandRoute.Put("/:brandID/update", brandConstructor.BrandHandler.Update)
	brandRoute.Delete("/:brandID/delete", brandConstructor.BrandHandler.Delete)

	// product routes

	productRoute := adminApi.Group("products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/:productID", productConstructor.ProductHandler.GetOne)
	productRoute.Post("/create", productConstructor.ProductHandler.Create)
	productRoute.Put(":productID/update", productConstructor.ProductHandler.Update)
	productRoute.Delete(":productID/delete", productConstructor.ProductHandler.Delete)

	// product-images routes

	pImageRoute := adminApi.Group("products/:productID/product-images")
	pImageRoute.Get("/", pImageConstructor.ProductImageHandler.GetAll)
	pImageRoute.Get("/:productImageID", pImageConstructor.ProductImageHandler.GetOne)
	pImageRoute.Post("/create", pImageConstructor.ProductImageHandler.Create)
	pImageRoute.Delete("/:productImageID/delete", pImageConstructor.ProductImageHandler.Delete)

	// slider routes
	sliderRoute := adminApi.Group("sliders")
	sliderRoute.Get("", sliderConstructor.SliderHandler.GetAll)
	sliderRoute.Get("/:sliderID", sliderConstructor.SliderHandler.GetOne)
	sliderRoute.Post("/create", sliderConstructor.SliderHandler.Create)
	sliderRoute.Put(":sliderID/update", sliderConstructor.SliderHandler.UpdateStatus)
	sliderRoute.Delete("/:sliderID/delete", sliderConstructor.SliderHandler.Delete)
}
