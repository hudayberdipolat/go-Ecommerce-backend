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

func (pImageRepo productImageRepoImp) FindAll(productID int) ([]models.ProductImage, error) {
	var productImages []models.ProductImage

	if err := pImageRepo.db.Where("product_id=?", productID).Find(&productImages).Error; err != nil {
		return nil, err
	}
	return productImages, nil
}

func (pImageRepo productImageRepoImp) FindOne(productID, productImageID int) (*models.ProductImage, error) {
	var productImage models.ProductImage
	if err := pImageRepo.db.Where("id=?", productImageID).Where("product_id=?", productID).First(&productImage).Error; err != nil {
		return nil, err
	}
	return &productImage, nil
}

func (pImageRepo productImageRepoImp) Store(productImage models.ProductImage) error {
	if err := pImageRepo.db.Create(&productImage).Error; err != nil {
		return err
	}
	return nil
}

func (pImageRepo productImageRepoImp) Destroy(productID, productImageID int) error {
	var productImage models.ProductImage
	if err := pImageRepo.db.Where("id=?", productImageID).Where("product_id=?", productID).Unscoped().Delete(&productImage).Error; err != nil {
		return err
	}
	return nil
}
