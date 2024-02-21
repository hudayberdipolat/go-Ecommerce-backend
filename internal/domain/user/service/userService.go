package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"

type UserService interface {
	Register(registerData dto.RegisterRequest) (*dto.AuthUserReponse, error)
	Login(loginData dto.LoginRequest) (*dto.AuthUserReponse, error)
	UpdateData(userID int, updateData dto.UpdateUserData) error
	UpdateUserPassword(userID int, updatePassword dto.UpdateUserPassword) error
}
