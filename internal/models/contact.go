package models

type Contact struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func (*Contact) TableName() string {
	return "contacts"
}
