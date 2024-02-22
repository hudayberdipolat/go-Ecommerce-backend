package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/service"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) UpdateData(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) UpdatePassword(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Login(ctx *fiber.Ctx) error {
	panic("admin handler imp")
}
