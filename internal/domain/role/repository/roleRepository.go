package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type RoleRepository interface {
	FindOne(roleID int) (*models.Role, error)
	FindAll() ([]models.Role, error)
	Create(role models.Role) error
	Update(roleID int, role models.Role) error
	Delete(roleID int) error
}
