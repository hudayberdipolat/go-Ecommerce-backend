package dto

type RoleRequest struct {
	RoleName string `json:"role_name" validate:"required"`
}
