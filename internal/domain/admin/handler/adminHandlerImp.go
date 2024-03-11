package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type adminHandlerImp struct {
	adminService service.AdminService
	config       *config.Config
}

func NewAdminHandler(service service.AdminService, config *config.Config) AdminHandler {
	return adminHandlerImp{
		adminService: service,
		config:       config,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}

func (adminHandler adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}

func (adminHandler adminHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}

func (adminHandler adminHandlerImp) UpdateData(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}

func (adminHandler adminHandlerImp) UpdatePassword(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}

func (adminHandler adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("admin Handler IMP")
}
