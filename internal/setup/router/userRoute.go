package router

import (
	"github.com/gofiber/fiber/v2"
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
	userRoute.Post("/update-profile", userConstructor.UserHandler.Update)
	userRoute.Post("/update-profile", userConstructor.UserHandler.ChangePassword)

}
