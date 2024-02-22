package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type GetOneAdminResponse struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	AdminStatus string `json:"admin_status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewGetOneAdminResponse(admin *models.Admin) GetOneAdminResponse {
	return GetOneAdminResponse{
		ID:          admin.ID,
		Firstname:   admin.Firstname,
		Surname:     admin.Surname,
		PhoneNumber: admin.PhoneNumber,
		AdminRole:   admin.AdminRole,
		AdminStatus: admin.AdminStatus,
		CreatedAt:   admin.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   admin.UpdatedAt.Format("01-02-2006"),
	}
}

type GetAllAdminResponse struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	AdminStatus string `json:"admin_status"`
}

func NewGetAllAdminResponse(admins []models.Admin) []GetAllAdminResponse {
	var allAdminResponse []GetAllAdminResponse

	for _, admin := range admins {
		adminResponse := GetAllAdminResponse{
			ID:          admin.ID,
			Firstname:   admin.Firstname,
			Surname:     admin.Surname,
			PhoneNumber: admin.PhoneNumber,
			AdminRole:   admin.AdminRole,
			AdminStatus: admin.AdminStatus,
		}

		allAdminResponse = append(allAdminResponse, adminResponse)
	}
	return allAdminResponse
}

type AdminAuthResponse struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	AdminStatus string `json:"admin_status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"acess_token"`
}

func NewAdminAuthResponse(admin *models.Admin, accessToken string) AdminAuthResponse {
	return AdminAuthResponse{
		ID:          admin.ID,
		Firstname:   admin.Firstname,
		Surname:     admin.Surname,
		PhoneNumber: admin.PhoneNumber,
		AdminRole:   admin.AdminRole,
		AdminStatus: admin.AdminStatus,
		CreatedAt:   admin.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   admin.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
