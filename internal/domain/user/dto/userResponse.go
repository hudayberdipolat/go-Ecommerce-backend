package dto

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type UserResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"access_token"`
}

func NewUserResponse(user *models.User, accessToken string) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
