package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepoImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepoImp{
		db: db,
	}
}

func (adminRepo adminRepoImp) FindAll() (*models.Admin, error) {
	panic("admin Repo IMP")
}

func (adminRepo adminRepoImp) FindOne(adminID int) (*models.Admin, error) {
	panic("admin Repo IMP")
}

func (adminRepo adminRepoImp) Store(admin models.Admin) error {
	panic("admin Repo IMP")
}

func (adminRepo adminRepoImp) Update(adminID int, updateAdmin models.Admin) error {
	panic("admin Repo IMP")
}

func (adminRepo adminRepoImp) UpdateAdminPassword(adminID int, password string) error {
	panic("admin Repo IMP")
}

func (adminRepo adminRepoImp) Destroy(adminID int) error {
	panic("admin Repo IMP")
}
