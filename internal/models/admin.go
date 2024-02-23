package models

import "time"

type Admin struct {
	ID          int          `json:"id"`
	Firstname   string       `json:"firstname"`
	Surname     string       `json:"surname"`
	PhoneNumber string       `json:"phone_number"`
	AdminRole   string       `json:"admin_role"`
	AdminStatus string       `json:"admin_status"`
	Password    string       `json:"password"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Role        []Role       `json:"roles"`
	Permissions []Permission `json:"permissions"`
}

func (*Admin) TableName() string {
	return "admins"
}
