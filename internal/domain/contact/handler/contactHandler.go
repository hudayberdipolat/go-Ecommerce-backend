package handler

import "github.com/gofiber/fiber/v2"

type ContactHandler interface {
	GetContact(ctx *fiber.Ctx) error
	UpdateContact(ctx *fiber.Ctx) error
}
