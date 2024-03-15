package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindOne(userID int) (*models.User, error)
	UpdateUserStatus(userID int, userStatus string) error
	FindUserByPhoneNumber(phoneNumber string) (*models.User, error)
	FindUserByID(userID int) (*models.User, error)
	Store(user models.User) error
	GetUser(userID int, phoneNumber string) (*models.User, error)
	Update(userID int, updateUser models.User) error
	ChangePassword(userID int, password string) error

	// for admins

}
