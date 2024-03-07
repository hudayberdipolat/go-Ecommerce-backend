package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"

type UserService interface {
	RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserAuthResponse, error)
	LoginUser(loginRequest dto.LoginRequest) (*dto.UserAuthResponse, error)
	//  USER DATA
	GetUser(userID int, phoneNumber string) (*dto.UserResponse, error)
	UpdateUser(userID int, request dto.UpdateUserRequest) error
	ChangeUserPassword(userID int, request dto.ChangeUserPasswordRequest) error
}
