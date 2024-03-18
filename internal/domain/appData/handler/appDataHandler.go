package handler

import "github.com/gofiber/fiber/v2"

type AppDataHandler interface {
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
