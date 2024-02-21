package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Surname     string    `json:"surname"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}
