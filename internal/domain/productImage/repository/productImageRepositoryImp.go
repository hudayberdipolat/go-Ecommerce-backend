package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type productImageRepositoryImp struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return productImageRepositoryImp{
		db: db,
	}
}

func (prodImageRepo productImageRepositoryImp) GetOne(productImage_id int) (*models.ProductImage, error) {

	panic("product image repo")
}

func (prodImageRepo productImageRepositoryImp) Create(image models.ProductImage) error {

	panic("product image repo")
}

func (prodImageRepo productImageRepositoryImp) Delete(productImage_id int) error {

	panic("product image repo")
}

func (prodImageRepo productImageRepositoryImp) GetOneProduct(productID int) (*models.Product, error) {

	panic("product image repo")
}
