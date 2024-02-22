package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		db: db,
	}
}

func (adminRepo adminRepositoryImp) FindAll() ([]models.Admin, error) {
	var admins []models.Admin

	if err := adminRepo.db.Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

func (adminRepo adminRepositoryImp) FindOneAdminWithID(adminID int) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.db.First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	if err := adminRepo.db.Create(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu telefon telgi eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (adminRepo adminRepositoryImp) UpdateAdminData(adminID int, admin models.Admin) error {
	if err := adminRepo.db.Model(&models.Admin{}).Where("id=?", adminID).Updates(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu telefon telgi eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (adminRepo adminRepositoryImp) UpdateAdminPassword(adminID int, newPassword string) error {
	if err := adminRepo.db.Model(&models.Admin{}).Where("id=?", adminID).Updates(models.Admin{
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	if err := adminRepo.db.Unscoped().Delete(&models.Admin{}, adminID).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepositoryImp) FindAdminWithPhoneNumber(phoneNumber string) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.db.Where("phone_number = ?", phoneNumber).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
