package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type roleServiceImp struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return roleServiceImp{
		roleRepo: repo,
	}
}

func (roleService roleServiceImp) GetOneRole(roleID int) (*dto.RoleResponse, error) {
	role, err := roleService.roleRepo.FindOne(roleID)
	if err != nil {
		return nil, err
	}
	roleResponse := dto.NewOneRoleResponse(role)
	return &roleResponse, nil
}

func (roleService roleServiceImp) GetAllRole() ([]dto.RoleResponse, error) {
	roles, err := roleService.roleRepo.FindAll()
	if err != nil {
		return nil, err
	}
	allRoleResponse := dto.NewAllRoleResponse(roles)
	return allRoleResponse, nil
}

func (roleService roleServiceImp) CreateRole(request dto.RoleRequest) error {
	createRole := models.Role{
		RoleName: request.RoleName,
	}

	if err := roleService.roleRepo.Create(createRole); err != nil {
		return err
	}
	return nil
}

func (roleService roleServiceImp) UpdateRole(roleID int, request dto.RoleRequest) error {
	updateRole, err := roleService.roleRepo.FindOne(roleID)
	if err != nil {
		return err
	}

	updateRole.RoleName = request.RoleName

	if err := roleService.roleRepo.Update(updateRole.ID, *updateRole); err != nil {
		return err
	}
	return nil
}

func (roleService roleServiceImp) DeleteRole(roleID int) error {
	deleteRole, err := roleService.roleRepo.FindOne(roleID)
	if err != nil {
		return err
	}

	if err := roleService.roleRepo.Delete(deleteRole.ID); err != nil {
		return nil
	}
	return nil
}
