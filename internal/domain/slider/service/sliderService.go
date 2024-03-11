package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type SliderService interface {
	GetAllSlider() ([]models.Slider, error)
	GetOneSlider(sliderID int) (*models.Slider, error)
	CreateSlider(createRequest dto.CreateSliderRequest) error
	UpdateSliderStatus(sliderID int, updateSlider models.Slider) error
	DeleteSlider(sliderID int) error
}
