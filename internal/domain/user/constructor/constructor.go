package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/user/service"
	"gorm.io/gorm"
)

var (
	userRepo    repository.UserRepository
	userService service.UserService
	UserHandler handler.UserHandler
)

func UserRequirmentCreator(db *gorm.DB) {
	userRepo = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
	UserHandler = handler.NewUserHandler(userService)
}
