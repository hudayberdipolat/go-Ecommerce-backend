package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserStatus  string    `json:"user_status"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}
