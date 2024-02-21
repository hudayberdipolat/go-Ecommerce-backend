package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AuthUserReponse struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"access_token"`
}

func NewAuthUserResponse(user models.User, accessToken string) AuthUserReponse {
	return AuthUserReponse{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format("1-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
