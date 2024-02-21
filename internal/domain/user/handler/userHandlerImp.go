package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/service"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) Register(ctx *fiber.Ctx) error {
	panic("user handler imp")
}

func (userHandler userHandlerImp) Login(ctx *fiber.Ctx) error {
	panic("user handler imp")
}
