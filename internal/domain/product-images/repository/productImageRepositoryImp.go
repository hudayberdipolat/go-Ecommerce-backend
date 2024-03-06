package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type productImageRepoImp struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return productImageRepoImp{
		db: db,
	}
}

func (pImageRepo productImageRepoImp) FindAll() ([]models.ProductImage, error) {
	panic("productImage repo imp")
}

func (pImageRepo productImageRepoImp) FindOne(productID, productImageID int) (*models.ProductImage, error) {
	panic("productImage repo imp")
}

func (pImageRepo productImageRepoImp) Store(productImage models.ProductImage) error {
	panic("productImage repo imp")
}

func (pImageRepo productImageRepoImp) Destroy(productID, productImageID int) error {
	panic("productImage repo imp")
}
