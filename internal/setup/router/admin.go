package router

import (
	"github.com/gofiber/fiber/v2"
	aboutConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/constructor"
	adminConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/constructor"
	BrendConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/constructor"
	contactConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/constructor"
	permissionConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/constructor"
	productConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/constructor"
	productImageConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/constructor"
	roleConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/constructor"
)

func AdminRoutes(app *fiber.App) {
	// admin routes api
	adminApi := app.Group("/api/admin")

	// admin  auth
	adminAuth := app.Group("/api/auth")
	adminAuth.Post("/login", adminConstructor.AdminHandler.Login)

	// super admin for admin routes
	adminRoute := adminApi.Group("admins")
	adminRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminRoute.Post("/create", adminConstructor.AdminHandler.Create)
	adminRoute.Put("/update-data", adminConstructor.AdminHandler.UpdateData)
	adminRoute.Put("/update-password", adminConstructor.AdminHandler.UpdatePassword)
	adminRoute.Delete("/:adminID/delete", adminConstructor.AdminHandler.Delete)

	// role routes
	roleRoute := adminApi.Group("roles")
	roleRoute.Get("/", roleConstructor.RoleHandler.GetAll)
	roleRoute.Get("/:roleID", roleConstructor.RoleHandler.GetOne)
	roleRoute.Post("/create", roleConstructor.RoleHandler.Create)
	roleRoute.Put("/:roleID/update", roleConstructor.RoleHandler.Update)
	roleRoute.Delete("/:roleID/delete", roleConstructor.RoleHandler.Delete)

	// permission routes

	permissionRoute := adminApi.Group("permissions")
	permissionRoute.Get("/", permissionConstructor.PermissionHandler.GetAll)
	permissionRoute.Get("/:permissionID", permissionConstructor.PermissionHandler.GetOne)
	permissionRoute.Post("/create", permissionConstructor.PermissionHandler.Create)
	permissionRoute.Put("/:permissionID/update", permissionConstructor.PermissionHandler.Update)
	permissionRoute.Delete("/:permissionID/delete", permissionConstructor.PermissionHandler.Delete)

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
	aboutRoute.Put("/:aboutID/update", aboutConstructor.AboutHandler.Update)

}
