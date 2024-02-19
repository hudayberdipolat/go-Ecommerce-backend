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

func (repo categoryRepositoryImp) OneCategory(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := repo.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo categoryRepositoryImp) AllCategory() ([]models.Category, error) {
	var categories []models.Category
	if err := repo.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo categoryRepositoryImp) Create(category models.Category) error {
	if err := repo.db.Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu category ady eýýäm ulanylýar !!")
		}
		return err
	}
	return nil
}

func (repo categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	if err := repo.db.Model(&models.Category{}).Where("id=?", categoryID).Updates(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu category ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (repo categoryRepositoryImp) Delete(categoryID int) error {
	if err := repo.db.Unscoped().Delete(models.Category{}, categoryID).Error; err != nil {
		return err
	}
	return nil
}
