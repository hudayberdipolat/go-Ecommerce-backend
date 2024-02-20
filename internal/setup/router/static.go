package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

func StaticFile(app *fiber.App, config *config.Config) {
	app.Static("/public", config.FolderConfig.PublicPath)
}
