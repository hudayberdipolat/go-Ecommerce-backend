package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AppDataRepository interface {
	FindAll() ([]models.AppData, error)
	FindOne(appDataID int) (*models.AppData, error)
	Create(appData models.AppData) error
	Update(appDataID int, appData models.AppData) error
	Delete(appDataID int) error
}
