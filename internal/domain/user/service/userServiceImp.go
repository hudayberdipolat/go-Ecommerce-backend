package service

import (
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/jwtToken/userToken"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserResponse, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.MaxCost)

	user := models.User{
		Username:    registerRequest.Username,
		PhoneNumber: registerRequest.PhoneNumber,
		UserStatus:  "ACTIVE",
		Password:    string(password),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := userService.userRepo.Store(user); err != nil {
		return nil, err
	}

	getUser, err := userService.userRepo.FindUserByPhoneNumber(user.PhoneNumber)
	if err != nil {
		return nil, err
	}

	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	if err != nil {
		return nil, err
	}
	userRespone := dto.NewUserResponse(getUser, accessToken)

	return &userRespone, nil
}

func (userService userServiceImp) LoginUser(loginRequest dto.LoginRequest) (*dto.UserResponse, error) {

	getUser, err := userService.userRepo.FindUserByPhoneNumber(loginRequest.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginRequest.Password)); err != nil {
		return nil, err
	}

	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	if err != nil {
		return nil, err
	}
	userRespone := dto.NewUserResponse(getUser, accessToken)

	return &userRespone, nil

}
