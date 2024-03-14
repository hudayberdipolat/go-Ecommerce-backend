package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type adminHandlerImp struct {
	adminService service.AdminService
	config       *config.Config
}

func NewAdminHandler(service service.AdminService, config *config.Config) AdminHandler {
	return adminHandlerImp{
		adminService: service,
		config:       config,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	admins, err := adminHandler.adminService.GetAllAdmin()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "admins not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "get all admins", admins)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (adminHandler adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	admin, err := adminHandler.adminService.GetOneAdmin(adminID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "admin not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "get one admin", admin)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (adminHandler adminHandlerImp) Create(ctx *fiber.Ctx) error {
	var createAdminRequest dto.CreateAdminRequest

	// body parser
	if err := ctx.BodyParser(&createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data

	if err := validate.ValidateStruct(createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate phone number

	if err := adminHandler.adminService.CreateAdmin(ctx, adminHandler.config, createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "created not admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(http.StatusCreated, "created admin successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)

}

func (adminHandler adminHandlerImp) UpdateData(ctx *fiber.Ctx) error {
	var updateAdminRequest dto.UpdateAdminRequest
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))

	// body parser
	if err := ctx.BodyParser(&updateAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := adminHandler.adminService.UpdateAdmin(ctx, adminHandler.config, adminID, updateAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "updated can't admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "updated admin successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (adminHandler adminHandlerImp) UpdatePassword(ctx *fiber.Ctx) error {
	var updateAdminPasswordRequest dto.ChangeAdminPasswordRequest
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))

	// body parser
	if err := ctx.BodyParser(&updateAdminPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate data
	if err := validate.ValidateStruct(updateAdminPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := adminHandler.adminService.UpdateAdminPassword(adminID, updateAdminPasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't update password", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "updated admin password successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (adminHandler adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))

	if err := adminHandler.adminService.DeleteAdmin(adminID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "admin deleted successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}
