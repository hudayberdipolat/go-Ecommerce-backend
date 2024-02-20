package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AboutRepository interface {
	GetOne(aboutID int) (*models.About, error)
	Update(aboutID int, about models.About) error
}
