package models

type Permission struct {
	ID             int    `json:"id"`
	PermissionName string `json:"permision_name"`
}

func (*Permission) TableName() string {
	return "permissions"
}
