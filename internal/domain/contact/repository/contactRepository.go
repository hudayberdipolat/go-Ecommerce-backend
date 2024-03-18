package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ContactRepository interface {
	GetContact(contactID int) (*models.Contact, error)
	Update(contactID int, contact models.Contact) error
}
