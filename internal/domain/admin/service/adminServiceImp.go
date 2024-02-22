package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (adminService adminServiceImp) GetAllAdmin() ([]dto.GetAllAdminResponse, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*dto.GetOneAdminResponse, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) CreateAdmin(request dto.CreateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) UpdateAdminData(adminID int, requet dto.UpdateAdminDataRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) UpdateAdminPassword(adminID int, requet dto.UpdateAdminPasswordRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) LoginAdmin(request dto.AdminLoginRequest) (dto.AdminAuthResponse, error) {
	panic("admin service imp")
}
