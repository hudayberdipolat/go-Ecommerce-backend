package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type AdminService interface {
	GetAllAdmin() ([]models.Admin, error)
	GetOneAdmin(adminID int) (*models.Admin, error)
	CreateAdmin(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateAdminRequest) error
	UpdateAdmin(ctx *fiber.Ctx, config *config.Config, adminID int, updateRequest dto.UpdateAdminRequest) error
	UpdateAdminPassword(adminID int, updatePassword dto.ChangeAdminPasswordRequest) error
	DeleteAdmin(adminID int) error
}
