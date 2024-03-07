package repository

import (
	"errors"

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
	var product models.Product
	if err := productRepo.db.Preload("Category").Preload("SubCategory").
		Preload("Brand").Preload("ProductImages").Preload("ProductComments").
		First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepo productRepositoryImp) GetAll() ([]models.Product, error) {
	var products []models.Product

	if err := productRepo.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepo productRepositoryImp) Store(product models.Product) error {

	if err := productRepo.db.Create(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("dublicate error: --> product slug")
		}
		return err
	}
	return nil
}

func (productRepo productRepositoryImp) Update(productID int, product models.Product) error {
	var updateProduct models.Product

	if err := productRepo.db.Model(&updateProduct).Where("id=?", productID).Updates(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("dublicate error: --> product slug")
		}
		return err
	}
	return nil
}

func (productRepo productRepositoryImp) Destroy(productID int) error {
	var deleteProduct models.Product
	if err := productRepo.db.Where("id=?", productID).Unscoped().Delete(&deleteProduct).Error; err != nil {
		return err
	}
	return nil
}

// FOR FRONT

func (productRepo productRepositoryImp) GetOneBySlug(productSlug string) (*models.Product, error) {
	var product models.Product
	if err := productRepo.db.Where("product_slug=?", productSlug).Preload("Category").
		Preload("SubCategory").Preload("Brand").Preload("ProductImages").Preload("ProductComments").First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
