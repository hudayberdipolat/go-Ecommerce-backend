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

func (p productRepositoryImp) FindOne(productID int) (*models.Product, error) {
	var product models.Product
	if err := p.db.Preload("Category").Preload("Brend").Preload("ProductImages").First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p productRepositoryImp) FindAll() ([]models.Product, error) {
	var products []models.Product

	if err := p.db.Preload("Category").Preload("Brend").Preload("ProductImages").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepositoryImp) Create(product models.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p productRepositoryImp) Update(productID int, product models.Product) error {
	if err := p.db.Model(&models.Product{}).Where("id=?", productID).Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p productRepositoryImp) Delete(productID int) error {
	if err := p.db.Unscoped().Delete(models.Product{}, productID).Error; err != nil {
		return err
	}
	return nil
}

func (p productRepositoryImp) FindOneProductWithSlug(productSlug string) (*models.Product, error) {
	var product models.Product
	err := p.db.Where("product_slug=?", productSlug).Preload("Category").
		Preload("Brend").Preload("ProductImages").First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
