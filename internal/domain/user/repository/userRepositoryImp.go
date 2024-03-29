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

func (userRepo userRepositoryImp) FindAll() ([]models.User, error) {
	var users []models.User

	if err := userRepo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepo userRepositoryImp) FindOne(userID int) (*models.User, error) {
	var user models.User
	if err := userRepo.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) UpdateUserStatus(userID int, userStatus string) error {
	var user models.User
	if err := userRepo.db.Model(&user).Where("id=?", userID).Updates(&models.User{
		UserStatus: userStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (userRepo userRepositoryImp) FindUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := userRepo.db.Where("phone_number=?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) FindUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := userRepo.db.Where("id=?", userID).First(&user).Error; err != nil {
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

func (userRepo userRepositoryImp) GetUser(userID int, phoneNumber string) (*models.User, error) {
	var user models.User
	if err := userRepo.db.Where("id=?", userID).Where("phone_number=?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) Update(userID int, updateUser models.User) error {
	var user models.User

	if err := userRepo.db.Model(&user).Where("id =?", userID).Updates(updateUser).Error; err != nil {
		return err
	}
	return nil
}

func (userRepo userRepositoryImp) ChangePassword(userID int, password string) error {
	var user models.User

	if err := userRepo.db.Model(&user).Where("id=?", userID).Updates(&models.User{
		Password: password,
	}).Error; err != nil {
		return err
	}
	return nil
}
