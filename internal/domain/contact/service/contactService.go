package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"

type ContactService interface {
	GetOneContact(contactID int) (dto.ContactResponse, error)
	UpdateContact(contactID int, request dto.ContactRequest) error
}
