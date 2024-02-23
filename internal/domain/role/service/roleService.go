package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/dto"

type RoleService interface {
	GetOneRole(roleID int) (*dto.RoleResponse, error)
	GetAllRole() ([]dto.RoleResponse, error)
	CreateRole(request dto.RoleRequest) error
	UpdateRole(roleID int, request dto.RoleRequest) error
	DeleteRole(roleID int) error
}
