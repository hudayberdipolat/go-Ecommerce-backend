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
	panic("about repo imp")
}

func (aboutRepo aboutRepositoryImp) Update(aboutID int, about models.About) error {
	panic("about repo imp")
}
