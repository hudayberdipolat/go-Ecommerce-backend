package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type contactRepositoryImp struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return contactRepositoryImp{
		db: db,
	}
}

func (contactRepo contactRepositoryImp) GetContact(contactID int) (*models.Contact, error) {
	panic("contactRepo IMP")
}

func (contactRepo contactRepositoryImp) Update(contactID int) error {
	panic("contactRepo IMP")
}
