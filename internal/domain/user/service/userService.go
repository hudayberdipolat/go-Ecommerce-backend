package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"

type UserService interface {
	RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserResponse, error)
	LoginUser(loginRequest dto.LoginRequest) (*dto.UserResponse, error)
}
