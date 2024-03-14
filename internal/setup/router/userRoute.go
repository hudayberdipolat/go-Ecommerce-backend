package router

import (
	"github.com/gofiber/fiber/v2"
	brandConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	userConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/constructor"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/middleware"
)

func UserRoutes(app *fiber.App) {

	frontApi := app.Group("api/front")

	authRoute := frontApi.Group("auth")
	authRoute.Post("/register", userConstructor.UserHandler.Register)
	authRoute.Post("/login", userConstructor.UserHandler.Login)

	// USER ROUTES
	userRoute := frontApi.Group("/user")
	userRoute.Use(middleware.UserMiddleware)
	userRoute.Get("/", userConstructor.UserHandler.GetUser)
	userRoute.Put("/update-profile", userConstructor.UserHandler.Update)
	userRoute.Put("/change-password", userConstructor.UserHandler.ChangePassword)

	// USER WRITE COMMENT FOR PRODUCT

	// FRONT FOR GET ALL CATEGORIES AND GET ONE CATEGORY AND THEIR ARE PRODUCTS

	categoryRoute := frontApi.Group("categories")

	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categorySlug", categoryConstructor.CategoryHandler.GetOneCategory)

	// GET ALL BRENDS AND GET ONE BREND AND THEIR ARE PRODUCTS
	brandRoute := frontApi.Group("brands")
	brandRoute.Get("/", brandConstructor.BrandHandler.GetAll)
	brandRoute.Get("/:brandSlug", brandConstructor.BrandHandler.GetOneBrand)

	// GET ALL PRODUCTS AND GET ONE PRODUCT AND THEIR COMMENTS

	productRoute := frontApi.Group("products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/:productSlug", productConstructor.ProductHandler.GetOneProduct)
}
