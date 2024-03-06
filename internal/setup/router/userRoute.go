package router

import (
	"github.com/gofiber/fiber/v2"
	userConstructor "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/constructor"
)

func UserRoutes(app *fiber.App) {

	frontApi := app.Group("api/front")

	authRoute := frontApi.Group("auth")
	authRoute.Post("/register", userConstructor.UserHandler.Register)
	authRoute.Post("/register", userConstructor.UserHandler.Login)

}
