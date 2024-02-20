package dto

type ContactRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" validate:"required"`
}
