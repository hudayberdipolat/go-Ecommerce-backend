package service

import (
	"errors"

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

func (contactService contactServiceImp) GetOneContact(contactID int) (*dto.ContactResponse, error) {
	contact, err := contactService.contactRepo.GetOne(contactID)
	if err != nil {
		return nil, err
	}

	contactResponse := dto.NewContactResponse(contact)
	return &contactResponse, nil
}

func (contactService contactServiceImp) UpdateContact(contactID int, request dto.ContactRequest) error {
	updateContact, err := contactService.contactRepo.GetOne(contactID)
	if err != nil {
		return errors.New("contact update error : not found contact")
	}

	updateContact.PhoneNumber = request.PhoneNumber
	updateContact.Email = request.Email
	updateContact.Address = request.Address

	if err := contactService.contactRepo.Update(updateContact.ID, updateContact); err != nil {
		return err
	}

	return nil
}
