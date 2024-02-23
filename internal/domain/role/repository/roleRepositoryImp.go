package repository

import (
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
	panic("role repo imp")
}

func (roleRepo roleRepositoryImp) FindAll() ([]models.Role, error) {
	panic("role repo imp")
}

func (roleRepo roleRepositoryImp) Create(role models.Role) error {
	panic("role repo imp")
}

func (roleRepo roleRepositoryImp) Update(roleID int, role models.Role) error {
	panic("role repo imp")
}

func (roleRepo roleRepositoryImp) Delete(roleID int) error {
	panic("role repo imp")
}
