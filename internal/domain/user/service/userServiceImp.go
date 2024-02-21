package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/repository"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) Register(registerData dto.RegisterRequest) (dto.AuthUserReponse, error) {
	panic("user service imp")
}

func (userService userServiceImp) Login(loginData dto.LoginRequest) (dto.AuthUserReponse, error) {
	panic("user service imp")
}
