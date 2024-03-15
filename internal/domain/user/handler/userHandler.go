package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error

	UpdateUserStatus(ctx *fiber.Ctx) error

	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	//USER PROFILE DATA
	GetUser(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
}
