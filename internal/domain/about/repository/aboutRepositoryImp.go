package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type aboutRepositoryImp struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return aboutRepositoryImp{
		db: db,
	}
}

func (aboutRepo aboutRepositoryImp) GetOne(aboutID int) (*models.About, error) {
	var about models.About

	if err := aboutRepo.db.First(&about, aboutID).Error; err != nil {
		return nil, err
	}
	return &about, nil
}

func (aboutRepo aboutRepositoryImp) Update(aboutID int, about models.About) error {
	if err := aboutRepo.db.Model(&models.About{}).Where("id=?", aboutID).Updates(&about).Error; err != nil {
		return err
	}
	return nil
}
