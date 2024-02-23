package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/repository"
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
	panic("role service imp")
}

func (roleService roleServiceImp) GetAllRole() ([]dto.RoleResponse, error) {
	panic("role service imp")
}

func (roleService roleServiceImp) CreateRole(request dto.RoleRequest) error {
	panic("role service imp")
}

func (roleService roleServiceImp) UpdateRole(request dto.RoleRequest) error {
	panic("role service imp")
}

func (roleService roleServiceImp) DeleteRole(roleID int) error {
	panic("role service imp")
}
