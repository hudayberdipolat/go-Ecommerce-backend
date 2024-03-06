package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type UserRepository interface {
	FindUserByPhoneNumber(phoneNumber string) (*models.User, error)
	Store(user models.User) error
}
