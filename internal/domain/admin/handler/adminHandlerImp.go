package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	admins, err := adminHandler.adminService.GetAllAdmin()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "admins not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all admins data", admins)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (adminHandler adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))

	admin, err := adminHandler.adminService.GetOneAdmin(adminID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "admin not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get admin data", admin)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (adminHandler adminHandlerImp) Create(ctx *fiber.Ctx) error {
	var createAdminRequest dto.CreateAdminRequest

	// body parser

	if err := ctx.BodyParser(&createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate data

	if err := validate.ValidateStruct(createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "valiate  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(createAdminRequest.PhoneNumber)

	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi!!!", "Nädogry telefon belgi!!!", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create admin
	if err := adminHandler.adminService.CreateAdmin(createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "create admin error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}

	// return response
	successResponse := response.Success(http.StatusOK, "admin created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (adminHandler adminHandlerImp) UpdateData(ctx *fiber.Ctx) error {
	var updateAdminDataRequest dto.UpdateAdminDataRequest
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	if err := ctx.BodyParser(&updateAdminDataRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate data

	if err := validate.ValidateStruct(updateAdminDataRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "valiate  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(updateAdminDataRequest.PhoneNumber)

	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi!!!", "Nädogry telefon belgi!!!", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update admin data

	if err := adminHandler.adminService.UpdateAdminData(adminID, updateAdminDataRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "update admin  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// return response

	successResponse := response.Success(http.StatusOK, "admin data updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (adminHandler adminHandlerImp) UpdatePassword(ctx *fiber.Ctx) error {
	var updateAdminPasswordRequet dto.UpdateAdminPasswordRequest
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	if err := ctx.BodyParser(&updateAdminPasswordRequet); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate data

	if err := validate.ValidateStruct(updateAdminPasswordRequet); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// update admin password
	if err := adminHandler.adminService.UpdateAdminPassword(adminID, updateAdminPasswordRequet); err != nil {
		errResponse := response.Error(http.StatusNotFound, "update admin password  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// return response data
	successResponse := response.Success(http.StatusOK, "admin password changed successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (adminHandler adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))

	if err := adminHandler.adminService.DeleteAdmin(adminID); err != nil {
		errResponse := response.Error(http.StatusNotFound, "delete admin error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// return response data
	successResponse := response.Success(http.StatusOK, "admin deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// admin login data

func (adminHandler adminHandlerImp) Login(ctx *fiber.Ctx) error {
	var adminLoginRequest dto.AdminLoginRequest

	if err := ctx.BodyParser(&adminLoginRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	// validate data

	if err := validate.ValidateStruct(adminLoginRequest); err != nil {
		errResponse := response.Error(http.StatusNotFound, "validate  error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// validate phoneNumber
	validatePhoneNumber := validate.PhoneNumberValidate(adminLoginRequest.PhoneNumber)

	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi!!!", "Nädogry telefon belgi!!!", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// otp data

	loginResponse, err := adminHandler.adminService.LoginAdmin(adminLoginRequest)

	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "something wrong", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}

	// return response data
	successResponse := response.Success(http.StatusOK, "admin login successfully", loginResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}
