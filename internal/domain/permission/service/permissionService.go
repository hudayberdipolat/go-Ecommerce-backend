package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/dto"

type PermissionService interface {
	GetOneRole(roleID int) (*dto.PermissionResponse, error)
	GetAllRole() ([]dto.PermissionResponse, error)
	CreateRole(request dto.PermisionRequest) error
	UpdateRole(roleID int, request dto.PermisionRequest) error
	DeleteRole(roleID int) error
}
