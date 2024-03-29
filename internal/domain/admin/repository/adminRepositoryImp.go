package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepoImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepoImp{
		db: db,
	}
}

func (adminRepo adminRepoImp) FindAll() ([]models.Admin, error) {
	var admins []models.Admin
	if err := adminRepo.db.
		Select("id", "username", "full_name", "phone_number", "email", "admin_status", "admin_role", "created_at", "updated_at").
		Where("admin_role =?", "admin").Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminRepo adminRepoImp) FindOne(adminID int) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.db.First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepoImp) Store(admin models.Admin) error {
	if err := adminRepo.db.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepoImp) Update(adminID int, updateAdmin models.Admin) error {
	var admin models.Admin
	if err := adminRepo.db.Model(&admin).Where("id=?", adminID).Updates(&updateAdmin).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepoImp) UpdateAdminPassword(adminID int, password string) error {
	var admin models.Admin
	if err := adminRepo.db.Model(&admin).Where("id=?", adminID).Updates(&models.Admin{
		Password: password,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepoImp) Destroy(adminID int) error {
	var admin models.Admin
	if err := adminRepo.db.Where("id=?", adminID).Unscoped().Delete(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (adminRepo adminRepoImp) FindAdminWithPhoneNumber(phoneNumber string) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.db.Where("phone_number=?", phoneNumber).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
