package service

import (
	"errors"
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

func (userService userServiceImp) FindAllUser() ([]models.User, error) {
	users, err := userService.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService userServiceImp) FindOneUser(userID int) (*models.User, error) {
	user, err := userService.userRepo.FindOne(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService userServiceImp) UpdateUserStatus(userID int, request dto.UpdateUserStatusRequest) error {

	getUser, err := userService.userRepo.FindOne(userID)
	if err != nil {
		return err
	}

	if err := userService.userRepo.UpdateUserStatus(getUser.ID, request.UserStatus); err != nil {
		return err
	}
	return nil
}

func (userService userServiceImp) RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserAuthResponse, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

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
	userRespone := dto.NewUserAuthResponse(getUser, accessToken)

	return &userRespone, nil
}

func (userService userServiceImp) LoginUser(loginRequest dto.LoginRequest) (*dto.UserAuthResponse, error) {

	getUser, err := userService.userRepo.FindUserByPhoneNumber(loginRequest.PhoneNumber)
	if err != nil {
		return nil, errors.New("phone number or password wrong")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginRequest.Password)); err != nil {
		return nil, errors.New("phone number or password wrong")
	}

	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	if err != nil {
		return nil, err
	}
	userRespone := dto.NewUserAuthResponse(getUser, accessToken)

	return &userRespone, nil

}

func (userService userServiceImp) GetUser(userID int, phoneNumber string) (*dto.UserResponse, error) {
	user, err := userService.userRepo.GetUser(userID, phoneNumber)
	if err != nil {
		return nil, err
	}

	userResponse := dto.NewUserResponse(user)
	return &userResponse, nil
}

func (userService userServiceImp) UpdateUser(userID int, request dto.UpdateUserRequest) error {
	updateUser, err := userService.userRepo.FindUserByID(userID)
	if err != nil {
		return err
	}

	// update user

	updateUser.Username = request.Username
	updateUser.PhoneNumber = request.PhoneNumber

	if err := userService.userRepo.Update(updateUser.ID, *updateUser); err != nil {
		return err
	}
	return nil
}

func (userService userServiceImp) ChangeUserPassword(userID int, request dto.ChangeUserPasswordRequest) error {

	getUser, err := userService.userRepo.FindUserByID(userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.OldPassword)); err != nil {
		return errors.New("old password wrong")
	}

	if err := userService.userRepo.ChangePassword(getUser.ID, request.Password); err != nil {
		return err
	}
	return nil
}
