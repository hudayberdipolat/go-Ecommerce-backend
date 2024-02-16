package models

import "time"

type User struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

func (*User) TableName() string {
	return "users"
}
