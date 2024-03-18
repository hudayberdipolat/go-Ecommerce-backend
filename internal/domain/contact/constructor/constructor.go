package constructor

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/handler"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/service"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/gorm"
)

var (
	contactRepo    repository.ContactRepository
	contactService service.ContactService
	ContactHandler handler.ContactHandler
)

func ContactRequirmentCreator(db *gorm.DB, config *config.Config) {
	contactRepo = repository.NewContactRepository(db)
	contactService = service.NewContactService(contactRepo)
	ContactHandler = handler.NewContactHandler(contactService, config)
}
