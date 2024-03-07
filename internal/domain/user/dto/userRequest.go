package dto

type RegisterRequest struct {
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Username        string `json:"username" validate:"required,min=3"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Username    string `json:"username" validate:"required,min=3"`
}

type ChangeUserPasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
