package handler

import (
	"net/http"

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

	// validate phone number

	validatePhoneNumber := validate.PhoneNumberValidate(registerRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	user, err := userHandler.userService.RegisterUser(registerRequest)

	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not register", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user register successfully", user)
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

	// validate phone number

	validatePhoneNumber := validate.PhoneNumberValidate(loginRequest.Password)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	user, err := userHandler.userService.LoginUser(loginRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not login", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user login successfully", user)
	return ctx.Status(http.StatusCreated).JSON(successResponse)

}

// USER PROFILE DATA

func (userHandler userHandlerImp) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	phoneNumber := ctx.Locals("phone_number").(string)
	user, err := userHandler.userService.GetUser(userID, phoneNumber)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not found something wrong...", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "user profile data", user)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

// UPDATE USER PROFILE DATA

func (userHandler userHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateUserRequest
	userID := ctx.Locals("user_id").(int)
	// body parser
	if err := ctx.BodyParser(&updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := userHandler.userService.UpdateUser(userID, updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated user profile data", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user profile updated successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)

}

// UPDATE USER PASSWORD

func (userHandler userHandlerImp) ChangePassword(ctx *fiber.Ctx) error {
	var updateRequest dto.ChangeUserPasswordRequest
	userID := ctx.Locals("user_id").(int)
	// body parser
	if err := ctx.BodyParser(&updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := userHandler.userService.ChangeUserPassword(userID, updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't change user password", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user password changed successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}
