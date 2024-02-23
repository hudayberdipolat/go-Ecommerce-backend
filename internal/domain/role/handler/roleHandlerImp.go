package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type roleHandlerImp struct {
	roleService service.RoleService
}

func NewRoleHandler(service service.RoleService) RoleHandler {
	return roleHandlerImp{
		roleService: service,
	}
}

func (roleHandler roleHandlerImp) GetOne(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))
	role, err := roleHandler.roleService.GetOneRole(roleID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "role not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get role data", role)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (roleHandler roleHandlerImp) GetAll(ctx *fiber.Ctx) error {
	role, err := roleHandler.roleService.GetAllRole()
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "roles not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all roles data", role)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (roleHandler roleHandlerImp) Create(ctx *fiber.Ctx) error {
	var roleRequet dto.RoleRequest

	if err := ctx.BodyParser(&roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := validate.ValidateStruct(roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// create role

	if err := roleHandler.roleService.CreateRole(roleRequet); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role created error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (roleHandler roleHandlerImp) Update(ctx *fiber.Ctx) error {
	var roleRequet dto.RoleRequest
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))

	if err := ctx.BodyParser(&roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := validate.ValidateStruct(roleRequet); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update role
	if err := roleHandler.roleService.UpdateRole(roleID, roleRequet); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role updated error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (roleHandler roleHandlerImp) Delete(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))
	if err := roleHandler.roleService.DeleteRole(roleID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role deleted error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
