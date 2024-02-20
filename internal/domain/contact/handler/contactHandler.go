package handler

import "github.com/gofiber/fiber/v2"

type ContactHandler interface {
	GetOne(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
