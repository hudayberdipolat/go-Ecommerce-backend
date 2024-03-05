package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type brandRepositoryImp struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return brandRepositoryImp{
		db: db,
	}
}

func (brandRepo brandRepositoryImp) GetOneByID(brandID int) (*models.Brand, error) {
	panic("brand repo imp")
}

func (brandRepo brandRepositoryImp) GetAll() ([]models.Brand, error) {
	panic("brand repo imp")
}

func (brandRepo brandRepositoryImp) Store(brand models.Brand) error {
	panic("brand repo imp")
}

func (brandRepo brandRepositoryImp) Update(brandID int, brand models.Brand) error {
	panic("brand repo imp")
}

func (brandRepo brandRepositoryImp) Destroy(brandID int) error {
	panic("brand repo imp")
}
