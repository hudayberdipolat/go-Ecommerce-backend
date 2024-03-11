package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type AdminRepository interface {
	FindAll() (*models.Admin, error)
	FindOne(adminID int) (*models.Admin, error)
	Store(admin models.Admin) error
	Update(adminID int, updateAdmin models.Admin) error
	UpdateAdminPassword(adminID int, password string) error
	Destroy(adminID int) error
}
