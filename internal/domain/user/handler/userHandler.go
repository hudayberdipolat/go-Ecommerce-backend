package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	UpdateUserData(ctx *fiber.Ctx) error
	UpdateUserPassword(ctx *fiber.Ctx) error
}
