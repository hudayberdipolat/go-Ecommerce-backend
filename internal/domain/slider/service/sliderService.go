package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type SliderService interface {
	GetAllSlider() ([]models.Slider, error)
	GetOneSlider(sliderID int) (*models.Slider, error)
	CreateSlider(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateSliderRequest) error
	UpdateSliderStatus(sliderID int, updateSlider dto.UpdateSliderStatus) error
	DeleteSlider(sliderID int) error
}
