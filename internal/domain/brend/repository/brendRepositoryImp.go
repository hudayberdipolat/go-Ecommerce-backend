package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type brendRepositoryImp struct {
	db *gorm.DB
}

func NewBrendRepository(db *gorm.DB) BrendRepository {
	return brendRepositoryImp{db: db}
}

func (b brendRepositoryImp) FindOneBrend(brendID int) (*models.Brend, error) {
	panic("brend repo imp")
}

func (b brendRepositoryImp) FindAllBrend(brendID int) (*models.Brend, error) {
	panic("brend repo imp")
}

func (b brendRepositoryImp) Create(brend models.Brend) error {
	panic("brend repo imp")
}

func (b brendRepositoryImp) Update(brendID int, brend models.Brend) error {
	panic("brend repo imp")
}

func (b brendRepositoryImp) Delete(brendID int) error {
	panic("brend repo imp")
}
