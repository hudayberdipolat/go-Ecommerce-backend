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
	var contact models.Contact
	if err := contactRepo.db.First(&contact, contactID).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (contactRepo contactRepositoryImp) Update(contactID int, contact *models.Contact) error {
	if err := contactRepo.db.Model(&models.Contact{}).Where("id=?", contactID).Updates(&contact).Error; err != nil {
		return err
	}
	return nil
}
