package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ContactResponse struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func NewContactResponse(contact *models.Contact) ContactResponse {
	return ContactResponse{
		ID:          contact.ID,
		PhoneNumber: contact.PhoneNumber,
		Email:       contact.Email,
		Address:     contact.Address,
	}
}
