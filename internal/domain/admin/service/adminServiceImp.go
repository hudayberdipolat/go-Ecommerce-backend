package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (adminServive adminServiceImp) GetAllAdmin() (*models.Admin, error) {
	panic("admin Service IMP")
}

func (adminServive adminServiceImp) GetOneAdmin(adminID int) (*models.Admin, error) {
	panic("admin Service IMP")
}

func (adminServive adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) (*models.Admin, error) {
	panic("admin Service IMP")
}

func (adminServive adminServiceImp) UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error {
	panic("admin Service IMP")
}

func (adminServive adminServiceImp) UpdateAdminPassword(adminID int, updatePassword dto.ChangeAdminPasswordRequest) error {
	panic("admin Service IMP")
}

func (adminServive adminServiceImp) DeleteAdmin(adminID int) error {
	panic("admin Service IMP")
}
