package repository

import (
	"errors"

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
	var brand models.Brand
	if err := brandRepo.db.First(&brand, brandID).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

func (brandRepo brandRepositoryImp) GetAll() ([]models.Brand, error) {
	var brands []models.Brand
	if err := brandRepo.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}

func (brandRepo brandRepositoryImp) Store(brand models.Brand) error {
	if err := brandRepo.db.Create(&brand).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("brand name already taken")
		}
		return err
	}

	return nil
}

func (brandRepo brandRepositoryImp) Update(brandID int, brand models.Brand) error {
	var updateBrand models.Brand
	if err := brandRepo.db.Model(&updateBrand).Where("id=?", brandID).Updates(&brand).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("brand name already taken")
		}
		return err
	}
	return nil
}

func (brandRepo brandRepositoryImp) Destroy(brandID int) error {
	var deleteBrand models.Brand
	if err := brandRepo.db.Unscoped().Delete(&deleteBrand, brandID).Error; err != nil {
		return err
	}
	return nil
}
