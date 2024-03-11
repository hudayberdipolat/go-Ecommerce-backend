package repository

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type SliderRepository interface {
	FindAll() ([]models.Slider, error)
	FindOne(sliderID int) (*models.Slider, error)
	Store(slider models.Slider) error
	UpdateStatus(sliderID int, updateSlider models.Slider) error
	Destroy(sliderID int) error
}
