package service

import (
	"errors"
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/contact/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
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
	contact, err := contactService.contactRepo.GetContact(contactID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (contactService contactServiceImp) UpdateContact(contactID int, updateRequest dto.UpdateContactRequest) error {
	// get contact
	contact, err := contactService.contactRepo.GetContact(contactID)
	if err != nil {
		return errors.New("contact not found")
	}

	contact.PhoneNumber = updateRequest.PhoneNumber
	contact.YouTubeAccount = updateRequest.YouTubeAccount
	contact.InstagramAccount = updateRequest.InstagramAccount
	contact.TiktokAccount = updateRequest.TiktokAccount
	contact.ImoAccount = updateRequest.ImoAccount
	contact.Address = updateRequest.Address
	contact.UpdatedAt = time.Now()

	if err := contactService.contactRepo.Update(contact.ID, *contact); err != nil {
		return err
	}
	return nil
}
