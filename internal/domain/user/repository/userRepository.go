package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type UserRepository interface {
	GetOneUserWithID(userID int) (*models.User, error)
	GetOneUserWithPhoneNumber(phoneNumber string) (*models.User, error)
	Create(user models.User) error
	UpdateData(userID int, user models.User) error
	UpdatePassword(userID int, password string) error
}