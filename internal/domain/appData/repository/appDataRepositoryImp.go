package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type appDataRepositoryImp struct {
	db *gorm.DB
}

func NewAppDataRepository(db *gorm.DB) AppDataRepository {
	return appDataRepositoryImp{
		db: db,
	}
}

func (appDataRepo appDataRepositoryImp) FindOne(appDataID int) (*models.AppData, error) {
	panic("appDataRepo IMP")
}

func (appDataRepo appDataRepositoryImp) Create(appData models.AppData) error {
	panic("appDataRepo IMP")
}

func (appDataRepo appDataRepositoryImp) Update(appDataID int, appData models.AppData) error {
	panic("appDataRepo IMP")
}
