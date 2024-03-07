package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetOne(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetOneProduct(ctx *fiber.Ctx) error
}
