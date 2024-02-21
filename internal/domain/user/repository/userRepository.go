package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type UserRepository interface {
	GetOne(phoneNumber string) (*models.User, error)
	Create(user models.User) error
	UpdateData(user models.User) error
	UpdatePassword(password string) error
}
