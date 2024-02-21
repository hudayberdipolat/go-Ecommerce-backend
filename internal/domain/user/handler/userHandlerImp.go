package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) Register(ctx *fiber.Ctx) error {

	var registerRequest dto.RegisterRequest

	// body parser
	if err := ctx.BodyParser(&registerRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate data
	if err := validate.ValidateStruct(registerRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(registerRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// otp validation with phoneNumber for register
	// otp bilen phone number hem tassyklanmaly
	// user register bolmazdan on validation bir gornusi

	// user register
	registerResponse, err := userHandler.userService.Register(registerRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "register user error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// return response
	successResponse := response.Success(http.StatusCreated, "user registered successfully", registerResponse)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (userHandler userHandlerImp) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest
	// body parser
	if err := ctx.BodyParser(&loginRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate data
	if err := validate.ValidateStruct(loginRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(loginRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// otp validation with phoneNumber for login
	// otp bilen phone number hem tassyklanmaly
	// user login bolmazdan on validation bir gornusi

	// login user

	loginResponse, err := userHandler.userService.Login(loginRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "register user error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// return response
	successResponse := response.Success(http.StatusCreated, "user login successfully", loginResponse)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (userHandler userHandlerImp) UpdateUserData(ctx *fiber.Ctx) error {
	var updateUserDataRequest dto.UpdateUserData
	userID, _ := strconv.Atoi(ctx.Params("userID"))
	// body parser
	if err := ctx.BodyParser(&updateUserDataRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate data
	if err := validate.ValidateStruct(updateUserDataRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(updateUserDataRequest.PhoneNumber)

	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update user data
	if err := userHandler.userService.UpdateData(userID, updateUserDataRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "update user profile data error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user profile data updated successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)

}

func (userHandler userHandlerImp) UpdateUserPassword(ctx *fiber.Ctx) error {
	var updateUserPasswordRequest dto.UpdateUserPassword
	userID, _ := strconv.Atoi(ctx.Params("userID"))
	// body parser
	if err := ctx.BodyParser(&updateUserPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate data
	if err := validate.ValidateStruct(updateUserPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update user password
	if err := userHandler.userService.UpdateUserPassword(userID, updateUserPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "update user password error error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "user password updated successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}
