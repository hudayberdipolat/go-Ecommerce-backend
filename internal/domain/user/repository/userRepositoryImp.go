package repository

import (
	"errors"

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

func (userRepo userRepositoryImp) GetOneUserWithID(userID int) (*models.User, error) {
	var user models.User
	if err := userRepo.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetOneUserWithPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User

	if err := userRepo.db.Where("phone_number=?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) Create(user models.User) error {
	if err := userRepo.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu telefon beglisi eýýäm ulanylýar")
		}
		return err
	}
	return nil
}

func (userRepo userRepositoryImp) UpdateData(userID int, user models.User) error {
	if err := userRepo.db.Model(&models.User{}).Where("id =?", userID).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu telefon beglisi eýýäm ulanylýar")
		}
		return err
	}
	return nil
}

func (userRepo userRepositoryImp) UpdatePassword(userID int, password string) error {
	if err := userRepo.db.Model(&models.User{}).Where("id=?", userID).Updates(&models.User{
		Password: password,
	}).Error; err != nil {
		return err
	}

	return nil
}
