package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type BrendRepository interface {
	FindOneBrend(brendID int) (*models.Brend, error)
	FindAllBrend() ([]models.Brend, error)
	Create(brend models.Brend) error
	Update(brendID int, brend models.Brend) error
	Delete(brendID int) error
	CheckBrendName(brendNameTk, brendNameRu string) bool
}
