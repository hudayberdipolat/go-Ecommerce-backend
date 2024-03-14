package dto

type CreateAdminRequest struct {
	Username        string `form:"username" validate:"required,min=3"`
	FullName        string `form:"full_name" validate:"required,min=3"`
	PhoneNumber     string `form:"phone_number" validate:"required,min=3"`
	Email           string `form:"email" validate:"required"`
	AdminImageURL   string `form:"admin_image_url"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateAdminRequest struct {
	Username      string `form:"username" validate:"required"`
	FullName      string `form:"full_name" validate:"required"`
	PhoneNumber   string `form:"phone_number" validate:"required"`
	Email         string `form:"email" validate:"required"`
	AdminImageURL string `form:"admin_image_url"`
	AdminStatys   string `form:"admin_status" validate:"required"`
}

type ChangeAdminPasswordRequest struct {
	OldPassword     string `form:"old_password" validate:"required"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginAdminRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
