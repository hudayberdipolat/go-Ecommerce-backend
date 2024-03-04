package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type categoryRepositoryImp struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepositoryImp{
		db: db,
	}
}

// admin for functions

func (categoryRepo categoryRepositoryImp) FindOneByID(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := categoryRepo.db.Preload("SubCategories").Preload("Products").
		First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (categoryRepo categoryRepositoryImp) FindAll() ([]models.Category, error) {
	var categories []models.Category
	if err := categoryRepo.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (categoryRepo categoryRepositoryImp) Store(category models.Category) error {

	tx := categoryRepo.db.Begin()
	err := tx.Create(&category).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("This category name already exists")
		}
		return err
	}
	tx.Commit()
	return nil
}

func (categoryRepo categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	var updateCategory models.Category
	if err := categoryRepo.db.Model(&updateCategory).Where("id=?", categoryID).Updates(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("This category name already exists")
		}
		return err
	}
	return nil
}

func (categoryRepo categoryRepositoryImp) Destroy(categoryID int) error {
	var deleteCategory models.Category
	if err := categoryRepo.db.Unscoped().Delete(&deleteCategory, categoryID).Error; err != nil {
		return err
	}
	return nil
}

// front for functions

func (categoryRepo categoryRepositoryImp) FindOneBySlug(categorySlug string) (*models.Category, error) {
	var category models.Category
	if err := categoryRepo.db.Where("category_slug=?", categorySlug).Preload("SubCategories").
		Preload("Products").First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
