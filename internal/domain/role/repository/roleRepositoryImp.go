package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type roleRepositoryImp struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepositoryImp{
		db: db,
	}
}

func (roleRepo roleRepositoryImp) FindOne(roleID int) (*models.Role, error) {
	var role models.Role
	if err := roleRepo.db.First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (roleRepo roleRepositoryImp) FindAll() ([]models.Role, error) {
	var roles []models.Role
	if err := roleRepo.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (roleRepo roleRepositoryImp) Create(role models.Role) error {
	if err := roleRepo.db.Create(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu role ady eýýäm bar!!!")
		}
		return err
	}
	return nil
}

func (roleRepo roleRepositoryImp) Update(roleID int, role models.Role) error {
	if err := roleRepo.db.Model(&models.Role{}).Where("id=?", roleID).Updates(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu role ady eýýäm bar!!!")
		}
		return err
	}
	return nil
}

func (roleRepo roleRepositoryImp) Delete(roleID int) error {
	var role models.Role
	if err := roleRepo.db.Unscoped().Delete(&role, roleID).Error; err != nil {
		return err
	}
	return nil
}
