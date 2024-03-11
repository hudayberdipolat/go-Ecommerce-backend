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
	panic("sliderRepo IMP")
}

func (sliderRepo sliderRepositoryImp) FindOne(sliderID int) (*models.Slider, error) {
	panic("sliderRepo IMP")
}

func (sliderRepo sliderRepositoryImp) Store(slider models.Slider) error {
	panic("sliderRepo IMP")
}

func (sliderRepo sliderRepositoryImp) UpdateStatus(sliderID int, updateSlider models.Slider) error {
	panic("sliderRepo IMP")
}

func (sliderRepo sliderRepositoryImp) Destroy(sliderID int) error {
	panic("sliderRepo IMP")
}
