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

func (prodImageRepo productImageRepositoryImp) GetOne(productImageID, productID int) (*models.ProductImage, error) {
	var productImage models.ProductImage
	err := prodImageRepo.db.Where("id=?", productImageID).Where("product_id=?", productID).First(&productImage).Error
	if err != nil {
		return nil, err
	}
	return &productImage, nil
}

func (prodImageRepo productImageRepositoryImp) Create(image models.ProductImage) error {
	if err := prodImageRepo.db.Create(image).Error; err != nil {
		return err
	}
	return nil
}

func (prodImageRepo productImageRepositoryImp) Delete(productImageID, productID int) error {
	var productImage models.ProductImage
	err := prodImageRepo.db.Where("id=?", productImageID).Where("product_id=?", productID).Unscoped().Delete(&productImage).Error
	if err != nil {
		return err
	}
	return nil
}

func (prodImageRepo productImageRepositoryImp) GetOneProduct(productID int) (*models.Product, error) {
	var product models.Product
	if err := prodImageRepo.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
