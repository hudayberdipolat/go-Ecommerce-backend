package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type PermissionRepository interface {
	FindOne(permissionID int) (*models.Permission, error)
	FindAll() ([]models.Permission, error)
	Create(permission models.Permission) error
	Update(permissionID int, Permission models.Permission) error
	Delete(permissionID int) error
}
