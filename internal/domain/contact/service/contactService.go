package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type ContactService interface {
	GetContact(contactID int) (*models.Contact, error)
	UpdateContact(contactID int, updateRequest dto.UpdateContactRequest) error
}
