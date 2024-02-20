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

func (contactRepo contactRepositoryImp) GetOne(contactID int) (*models.Contact, error) {
	panic("contact repo imp")
}

func (contactRepo contactRepositoryImp) Update(contactID, contact models.Contact) error {
	panic("contact repo imp")
}
