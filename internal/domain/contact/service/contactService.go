package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ContactService interface {
	GetContact(contactID int) (*models.Contact, error)
	UpdateContact(ctx *fiber.Ctx, config *config.Config, contactID int, updateRequest dto.UpdateContactRequest) error
}
