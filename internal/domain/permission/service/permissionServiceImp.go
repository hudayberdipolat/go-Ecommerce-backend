package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type permissionServiceImp struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return permissionServiceImp{
		permissionRepo: repo,
	}
}

func (permissionService permissionServiceImp) GetOneRole(roleID int) (*dto.PermissionResponse, error) {
	permission, err := permissionService.permissionRepo.FindOne(roleID)
	if err != nil {
		return nil, err
	}
	PermissionResponse := dto.NewOnePermissionResponse(permission)
	return &PermissionResponse, nil
}

func (permissionService permissionServiceImp) GetAllRole() ([]dto.PermissionResponse, error) {
	permissions, err := permissionService.permissionRepo.FindAll()
	if err != nil {
		return nil, err
	}
	allRoleResponse := dto.NewAllPermissionResponse(permissions)
	return allRoleResponse, nil
}

func (permissionService permissionServiceImp) CreateRole(request dto.PermisionRequest) error {

	createPermission := models.Permission{
		PermissionName: request.PermissionName,
	}

	if err := permissionService.permissionRepo.Create(createPermission); err != nil {
		return err
	}
	return nil
}

func (permissionService permissionServiceImp) UpdateRole(permissionID int, request dto.PermisionRequest) error {
	updatePermission, err := permissionService.permissionRepo.FindOne(permissionID)
	if err != nil {
		return err
	}

	updatePermission.PermissionName = request.PermissionName

	if err := permissionService.permissionRepo.Update(updatePermission.ID, *updatePermission); err != nil {
		return err
	}
	return nil
}

func (permissionService permissionServiceImp) DeleteRole(permissionID int) error {
	deletePermisssion, err := permissionService.permissionRepo.FindOne(permissionID)
	if err != nil {
		return err
	}

	if err := permissionService.permissionRepo.Delete(deletePermisssion.ID); err != nil {
		return nil
	}
	return nil
}
