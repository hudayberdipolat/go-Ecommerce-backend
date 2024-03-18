package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AppDataRepository interface {
	FindOne(appDataID int) (*models.AppData, error)
	Create(appData models.AppData) error
	Update(appDataID int, appData models.AppData) error
}
