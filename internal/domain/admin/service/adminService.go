package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"

type AdminService interface {
	GetAllAdmin() ([]dto.GetAllAdminResponse, error)
	GetOneAdmin(adminID int) (*dto.GetOneAdminResponse, error)
	CreateAdmin(request dto.CreateAdminRequest) error
	UpdateAdminData(adminID int, requet dto.UpdateAdminDataRequest) error
	UpdateAdminPassword(adminID int, requet dto.UpdateAdminPasswordRequest) error
	DeleteAdmin(adminID int) error
	//
	LoginAdmin(request dto.AdminLoginRequest) (dto.AdminAuthResponse, error)
}
