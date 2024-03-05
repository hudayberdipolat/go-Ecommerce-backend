package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type productRepositoryImp struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepositoryImp{
		db: db,
	}
}

func (productRepo productRepositoryImp) GetOneByID(productID int) (*models.Product, error) {
	panic("product Repo imp")
}

func (productRepo productRepositoryImp) GetAll() ([]models.Product, error) {
	panic("product Repo imp")
}

func (productRepo productRepositoryImp) Store(product models.Product) error {
	panic("product Repo imp")
}

func (productRepo productRepositoryImp) Update(productID int, product models.Product) error {
	panic("product Repo imp")
}

func (productRepo productRepositoryImp) Destroy(productID int) error {
	panic("product Repo imp")
}
