package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type permissionRepositoryImp struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return permissionRepositoryImp{
		db: db,
	}
}

func (permissionRepo permissionRepositoryImp) FindOne(permissionID int) (*models.Permission, error) {
	var permission models.Permission
	if err := permissionRepo.db.First(&permission, permissionID).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (permissionRepo permissionRepositoryImp) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := permissionRepo.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (permissionRepo permissionRepositoryImp) Create(role models.Permission) error {
	if err := permissionRepo.db.Create(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu permission ady eýýäm bar!!!")
		}
		return err
	}
	return nil
}

func (permissionRepo permissionRepositoryImp) Update(permissionID int, role models.Permission) error {
	if err := permissionRepo.db.Model(&models.Permission{}).Where("id=?", permissionID).Updates(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu permission ady eýýäm bar!!!")
		}
		return err
	}
	return nil
}

func (permissionRepo permissionRepositoryImp) Delete(permissionID int) error {
	var role models.Permission
	if err := permissionRepo.db.Unscoped().Delete(&role, permissionID).Error; err != nil {
		return err
	}
	return nil
}
