package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AdminAuthResponse struct {
	ID            int     `json:"id"`
	Username      string  `json:"username"`
	FullName      string  `json:"full_name"`
	PhoneNumber   string  `json:"phone_number"`
	Email         string  `jsin:"email"`
	AdminImageURL *string `json:"user_image_url"`
	AdminStatus   string  `json:"admin_status"`
	CrearedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	AccessToken   string  `json:"access_token"`
}

func NewAdminAuthResponse(admin *models.Admin, accessToken string) AdminAuthResponse {
	return AdminAuthResponse{
		ID:            admin.ID,
		Username:      admin.Username,
		FullName:      admin.PhoneNumber,
		Email:         admin.Email,
		AdminImageURL: admin.AdminImageURL,
		AdminStatus:   admin.AdminStatus,
		CrearedAt:     admin.CrearedAt.Format("01-02-2006"),
		UpdatedAt:     admin.UpdatedAt.Format("01-02-2006"),
		AccessToken:   accessToken,
	}
}
