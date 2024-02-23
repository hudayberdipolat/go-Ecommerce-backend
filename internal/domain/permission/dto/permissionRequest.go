package dto

type PermisionRequest struct {
	PermissionName string `json:"permission_name" validate:"required"`
}
