package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type subcategoryRepositoryImp struct {
	db *gorm.DB
}

func NewSubCategoryRespository(db *gorm.DB) SubCategoryRepository {
	return subcategoryRepositoryImp{
		db: db,
	}
}

func (subCategoryRepo subcategoryRepositoryImp) FindOne(subCategoryID int) (*models.SubCategory, error) {
	var subCategory models.SubCategory
	if err := subCategoryRepo.db.First(&subCategory, subCategoryID).Error; err != nil {
		return nil, err
	}

	return &subCategory, nil
}

func (subCategoryRepo subcategoryRepositoryImp) FindAll() ([]models.SubCategory, error) {
	var subCategories []models.SubCategory
	if err := subCategoryRepo.db.Find(&subCategories).Error; err != nil {
		return nil, err
	}
	return subCategories, nil
}

func (subCategoryRepo subcategoryRepositoryImp) Store(createRequest models.SubCategory) error {
	if err := subCategoryRepo.db.Create(&createRequest).Error; err != nil {
		return err
	}
	return nil
}

func (subCategoryRepo subcategoryRepositoryImp) Update(subCategoryID int, updateRequest models.SubCategory) error {
	var updateSubCategory models.SubCategory
	err := subCategoryRepo.db.Model(&updateSubCategory).Where("id=?", subCategoryID).Updates(&updateRequest).Error
	if err != nil {
		return err
	}
	return nil
}

func (subCategoryRepo subcategoryRepositoryImp) Destroy(categoryID, subCategoryID int) error {
	var deleteSubCategory models.SubCategory
	err := subCategoryRepo.db.Unscoped().Where("id=?", subCategoryID).Where("category_id=?", categoryID).Delete(&deleteSubCategory).Error
	if err != nil {
		return err
	}
	return nil
}
