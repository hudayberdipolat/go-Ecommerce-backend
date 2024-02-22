package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/service"
	"gorm.io/gorm"
)

var (
	adminRepo    repository.AdminRepository
	adminService service.AdminService
	AdminHandler handler.AdminHandler
)

func AdminRequirementsCreator(db *gorm.DB) {
	adminRepo = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepo)
	AdminHandler = handler.NewAdminHandler(adminService)
}
