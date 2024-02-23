package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/service"
)

type roleHandlerImp struct {
	roleService service.RoleService
}

func NewRoleHandler(service service.RoleService) RoleHandler {
	return roleHandlerImp{
		roleService: service,
	}
}

func (roleHandler roleHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("role handler imp")
}

func (roleHandler roleHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("role handler imp")
}

func (roleHandler roleHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("role handler imp")
}

func (roleHandler roleHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("role handler imp")
}

func (roleHandler roleHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("role handler imp")
}
