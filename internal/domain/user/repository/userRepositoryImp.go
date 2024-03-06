package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type userRepositoryImp struct {
	db gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: *db,
	}
}

func (userRepo userRepositoryImp) FindUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := userRepo.db.Where("phone_number=?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) Store(user models.User) error {
	if err := userRepo.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("this phoneNumber allready taken")
		}
		return err
	}
	return nil
}
