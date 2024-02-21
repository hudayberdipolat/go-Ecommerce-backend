package dto

type RegisterRequest struct {
	Firstname       string `json:"firstname" validate:"required"`
	Surname         string `json:"surname" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
