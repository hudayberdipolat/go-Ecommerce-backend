package dto

type AdminLoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type CreateAdminRequest struct {
	Firstname       string `json:"firstname" validate:"requried"`
	Surname         string `json:"surname" validate:"requried"`
	PhoneNumber     string `json:"phone_number" validate:"requried"`
	AdminRole       string `json:"admin_role" validate:"requried"`
	Password        string `json:"password" validate:"requried"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateAdminDataRequest struct {
	Firstname   string `json:"firstname" validate:"requried"`
	Surname     string `json:"surname" validate:"requried"`
	PhoneNumber string `json:"phone_number" validate:"requried"`
	AdminRole   string `json:"admin_role" validate:"requried"`
	AdminStatus string `json:"admin_status" validate:"requried"`
}

type UpdateAdminPasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"requried"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
