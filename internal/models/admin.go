package models

import "time"

type Admin struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	FullName      string    `json:"full_name"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `jsin:"email"`
	AdminImageURL *string   `json:"admin_image_url"`
	AdminStatus   string    `json:"admin_status"`
	AdminRole     string    `json:"admin_role"`
	Password      string    `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (*Admin) TableName() string {
	return "admins"
}
