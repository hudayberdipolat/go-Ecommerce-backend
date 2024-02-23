package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type PermissionResponse struct {
	ID             int    `json:"id"`
	PermissionName string `json:"permission_name"`
}

func NewOneRoleResponse(permission *models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:             permission.ID,
		PermissionName: permission.PermissionName,
	}
}

func NewAllPermissionResponse(permissions []models.Permission) []PermissionResponse {
	var allPermissionResponses []PermissionResponse
	for _, permission := range permissions {
		permissionResponse := PermissionResponse{
			ID:             permission.ID,
			PermissionName: permission.PermissionName,
		}
		allPermissionResponses = append(allPermissionResponses, permissionResponse)
	}
	return allPermissionResponses
}
