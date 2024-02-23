package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type RoleResponse struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

func NewOneRoleResponse(role models.Role) RoleResponse {
	return RoleResponse{
		ID:       role.ID,
		RoleName: role.RoleName,
	}
}

func NewAllRoleResponse(roles []models.Role) []RoleResponse {
	var allRoleResponses []RoleResponse
	for _, role := range roles {
		roleResponse := NewOneRoleResponse(role)
		allRoleResponses = append(allRoleResponses, roleResponse)
	}
	return allRoleResponses
}
