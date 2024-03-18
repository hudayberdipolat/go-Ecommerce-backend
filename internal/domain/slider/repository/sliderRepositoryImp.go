package repository

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type sliderRepositoryImp struct {
	db *gorm.DB
}

func NewSliderRepository(db *gorm.DB) SliderRepository {
	return sliderRepositoryImp{
		db: db,
	}
}

func (sliderRepo sliderRepositoryImp) FindAll() ([]models.Slider, error) {
	var sliders []models.Slider
	if err := sliderRepo.db.Find(&sliders).Error; err != nil {
		return nil, err
	}
	return sliders, nil
}

func (sliderRepo sliderRepositoryImp) FindOne(sliderID int) (*models.Slider, error) {
	var slider models.Slider
	if err := sliderRepo.db.First(&slider, sliderID).Error; err != nil {
		return nil, err
	}
	return &slider, nil
}

func (sliderRepo sliderRepositoryImp) Store(slider models.Slider) error {
	if err := sliderRepo.db.Create(&slider).Error; err != nil {
		return err
	}
	return nil
}

func (sliderRepo sliderRepositoryImp) UpdateStatus(sliderID int, updateSlider models.Slider) error {
	var slider models.Slider
	if err := sliderRepo.db.Model(&slider).Where("id=?", sliderID).Updates(&updateSlider).Error; err != nil {
		return err
	}
	return nil
}

func (sliderRepo sliderRepositoryImp) Destroy(sliderID int) error {
	var slider models.Slider
	if err := sliderRepo.db.Where("id=?", sliderID).Unscoped().Delete(&slider).Error; err != nil {
		return err
	}
	return nil
}
