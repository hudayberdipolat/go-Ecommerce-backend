package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ContactRepository interface {
	GetOne(contactID int) (*models.Contact, error)
	Update(contactID, contact models.Contact) error
}
