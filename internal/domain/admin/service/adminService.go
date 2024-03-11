package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type AdminService interface {
	GetAllAdmin() (*models.Admin, error)
	GetOneAdmin(adminID int) (*models.Admin, error)
	CreateAdmin(createRequest dto.CreateAdminRequest) (*models.Admin, error)
	UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error
	UpdateAdminPassword(adminID int, updatePassword dto.ChangeAdminPasswordRequest) error
	DeleteAdmin(adminID int) error
}
