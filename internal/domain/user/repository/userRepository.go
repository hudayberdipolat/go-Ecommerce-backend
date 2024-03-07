package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type UserRepository interface {
	FindUserByPhoneNumber(phoneNumber string) (*models.User, error)
	FindUserByID(userID int) (*models.User, error)
	Store(user models.User) error
	GetUser(userID int, phoneNumber string) (*models.User, error)
	Update(userID int, updateUser models.User) error
	ChangePassword(userID int, password string) error
}
