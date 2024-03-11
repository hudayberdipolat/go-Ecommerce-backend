package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type sliderServiceImp struct {
	sliderRepo repository.SliderRepository
}

func NewSliderService(repo repository.SliderRepository) SliderService {
	return sliderServiceImp{
		sliderRepo: repo,
	}
}

func (sliderService sliderServiceImp) GetAllSlider() ([]models.Slider, error) {
	panic("slider Service IMP")
}

func (sliderService sliderServiceImp) GetOneSlider(sliderID int) (*models.Slider, error) {
	panic("slider Service IMP")
}

func (sliderService sliderServiceImp) CreateSlider(createRequest dto.CreateSliderRequest) error {
	panic("slider Service IMP")
}

func (sliderService sliderServiceImp) UpdateSliderStatus(sliderID int, updateSlider models.Slider) error {
	panic("slider Service IMP")
}

func (sliderService sliderServiceImp) DeleteSlider(sliderID int) error {
	panic("slider Service IMP")
}
