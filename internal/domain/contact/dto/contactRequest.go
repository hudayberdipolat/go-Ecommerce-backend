package dto

type ContactRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" validate:"required"`
}
