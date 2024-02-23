package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type permissionHandlerImp struct {
	permissionService service.PermissionService
}

func NewRoleHandler(service service.PermissionService) PermissionHandler {
	return permissionHandlerImp{
		permissionService: service,
	}
}

func (permissionHandler permissionHandlerImp) GetOne(ctx *fiber.Ctx) error {
	permissionID, _ := strconv.Atoi(ctx.Params("permissionID"))
	role, err := permissionHandler.permissionService.GetOneRole(permissionID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "permission not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get permission data", role)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (permissionHandler permissionHandlerImp) GetAll(ctx *fiber.Ctx) error {
	role, err := permissionHandler.permissionService.GetAllRole()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "permissions not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all permissions data", role)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (permissionHandler permissionHandlerImp) Create(ctx *fiber.Ctx) error {
	var roleRequet dto.PermisionRequest

	if err := ctx.BodyParser(&roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := validate.ValidateStruct(roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// create role

	if err := permissionHandler.permissionService.CreateRole(roleRequet); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "permission created error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "permission created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (permissionHandler permissionHandlerImp) Update(ctx *fiber.Ctx) error {
	var roleRequet dto.PermisionRequest
	permissionID, _ := strconv.Atoi(ctx.Params("permissionID"))

	if err := ctx.BodyParser(&roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := validate.ValidateStruct(roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update role
	if err := permissionHandler.permissionService.UpdateRole(permissionID, roleRequet); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "permission updated error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "permission updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (permissionHandler permissionHandlerImp) Delete(ctx *fiber.Ctx) error {
	permissionID, _ := strconv.Atoi(ctx.Params("permissionID"))
	if err := permissionHandler.permissionService.DeleteRole(permissionID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role deleted error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "permission deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
