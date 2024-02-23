package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/permission/service"
	"gorm.io/gorm"
)

var (
	permissionRepo    repository.PermissionRepository
	permissionService service.PermissionService
	PermissionHandler handler.PermissionHandler
)

func PermissionRequirementsCreator(db *gorm.DB) {
	permissionRepo = repository.NewPermissionRepository(db)
	permissionService = service.NewPermissionService(permissionRepo)
	PermissionHandler = handler.NewPermissionHandler(permissionService)
}
