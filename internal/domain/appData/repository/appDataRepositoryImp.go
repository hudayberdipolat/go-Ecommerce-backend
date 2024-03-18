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
	var appData models.AppData

	if err := appDataRepo.db.First(&appData, appDataID).Error; err != nil {
		return nil, err
	}
	return &appData, nil
}

func (appDataRepo appDataRepositoryImp) Create(appData models.AppData) error {
	if err := appDataRepo.db.Create(&appData).Error; err != nil {
		return err
	}
	return nil
}

func (appDataRepo appDataRepositoryImp) Update(appDataID int, appData models.AppData) error {
	var updateAppData models.AppData
	err := appDataRepo.db.Model(&updateAppData).Where("id=?", appDataID).Updates(&appData).Error
	if err != nil {
		return err
	}
	return nil
}
