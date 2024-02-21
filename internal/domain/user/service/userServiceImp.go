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

func (userService userServiceImp) Register(registerData dto.RegisterRequest) (*dto.AuthUserReponse, error) {

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.MaxCost)

	user := models.User{
		Firstname:   registerData.Firstname,
		Surname:     registerData.Surname,
		UserStatus:  "active",
		PhoneNumber: registerData.PhoneNumber,
		Password:    string(passwordHash),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	//	panic("user service imp")

	if err := userService.userRepo.Create(user); err != nil {
		return nil, err
	}

	getUser, err := userService.userRepo.GetOneUserWithPhoneNumber(user.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// generate token

	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)

	if err != nil {
		return nil, err
	}

	registerResponse := dto.NewAuthUserResponse(getUser, accessToken)
	return &registerResponse, nil
}

func (userService userServiceImp) Login(loginData dto.LoginRequest) (*dto.AuthUserReponse, error) {
	getUser, err := userService.userRepo.GetOneUserWithPhoneNumber(loginData.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginData.Password)); err != nil {
		return nil, errors.New("Nädogry telefon belgi ýa-da password")
	}

	// generate token

	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)

	if err != nil {
		return nil, err
	}

	loginResponse := dto.NewAuthUserResponse(getUser, accessToken)
	return &loginResponse, nil

}

func (userService userServiceImp) UpdateData(userID int, updateData dto.UpdateUserData) error {
	updateUser, err := userService.userRepo.GetOneUserWithID(userID)
	if err != nil {
		return err
	}
	updateUser.Firstname = updateData.Firstname
	updateUser.Surname = updateData.Surname
	updateUser.PhoneNumber = updateData.PhoneNumber

	if err := userService.userRepo.UpdateData(updateUser.ID, *updateUser); err != nil {
		return err
	}
	return nil
}

func (userService userServiceImp) UpdateUserPassword(userID int, updatePassword dto.UpdateUserPassword) error {
	updateUserPassword, err := userService.userRepo.GetOneUserWithID(userID)
	if err != nil {
		return err
	}

	// old password barlag
	if err := bcrypt.CompareHashAndPassword([]byte(updateUserPassword.Password), []byte(updatePassword.OldPassword)); err != nil {
		return errors.New("Köne password nädogry!!!")
	}

	if err := userService.userRepo.UpdatePassword(updateUserPassword.ID, updatePassword.Password); err != nil {
		return err
	}
	return nil
}
