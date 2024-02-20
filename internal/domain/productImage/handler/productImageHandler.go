package handler

import "github.com/gofiber/fiber/v2"

type ProductImageHandler interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
