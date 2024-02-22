package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		db: db,
	}
}

func (adminRepo adminRepositoryImp) FindAll() ([]models.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) FindOneAdminWithID(adminID int) (*models.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) UpdateAdminData(adminID int, admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) UpdateAdminPassword(adminID int, admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	panic("admin repo imp")
}

//

func (adminRepo adminRepositoryImp) FindAdminWithPhoneNumber(phoneNumber string) (*models.Admin, error) {
	panic("admin repo imp")
}
