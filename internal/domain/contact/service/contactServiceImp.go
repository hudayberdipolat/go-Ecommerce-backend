package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type contactServiceImp struct {
	contactRepo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return contactServiceImp{
		contactRepo: repo,
	}
}

func (contactService contactServiceImp) GetContact(contactID int) (*models.Contact, error) {
	panic("contact Service IMP")
}

func (contactService contactServiceImp) UpdateContact(ctx *fiber.Ctx, config *config.Config, contactID int, updateRequest dto.UpdateContactRequest) error {
	panic("contact Service IMP")
}
