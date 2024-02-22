package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AdminRepository interface {
	FindAll() ([]models.Admin, error)
	FindOneAdminWithID(adminID int) (*models.Admin, error)
	Create(admin models.Admin) error
	UpdateAdminData(adminID int, admin models.Admin) error
	UpdateAdminPassword(adminID int, admin models.Admin) error
	Delete(adminID int) error

	//
	FindAdminWithPhoneNumber(phoneNumber string) (*models.Admin, error)
}
