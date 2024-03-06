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

	user, err := userHandler.userService.LoginUser(loginRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not login", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "user login successfully", user)
	return ctx.Status(http.StatusCreated).JSON(successResponse)

}
