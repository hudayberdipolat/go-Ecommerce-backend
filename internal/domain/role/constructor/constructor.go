package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/role/service"
	"gorm.io/gorm"
)

var (
	roleRepo    repository.RoleRepository
	roleService service.RoleService
	RoleHandler handler.RoleHandler
)

func RoleRequirementsCreator(db *gorm.DB) {
	roleRepo = repository.NewRoleRepository(db)
	roleService = service.NewRoleService(roleRepo)
	RoleHandler = handler.NewRoleHandler(roleService)
}
