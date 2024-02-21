package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (userRepo userRepositoryImp) GetOne(phoneNumber string) (*models.User, error) {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) Create(user models.User) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) UpdateData(user models.User) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) UpdatePassword(password string) error {
	panic("user repo imp")
}
