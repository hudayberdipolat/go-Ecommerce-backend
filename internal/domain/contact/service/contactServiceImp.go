package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/repository"
)

type contactServiceImp struct {
	contactRepo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return contactServiceImp{
		contactRepo: repo,
	}
}

func (contactService contactServiceImp) GetOneContact(contactID int) (dto.ContactResponse, error) {
	panic("contact serive imp")
}

func (contactService contactServiceImp) UpdateContact(contactID int, request dto.ContactRequest) error {

	panic("contact serive imp")
}
