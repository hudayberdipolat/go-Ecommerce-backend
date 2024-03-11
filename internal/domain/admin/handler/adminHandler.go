package handler

import "github.com/gofiber/fiber/v2"

type AdminHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	UpdateData(ctx *fiber.Ctx) error
	UpdatePassword(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
