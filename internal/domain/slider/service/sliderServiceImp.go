package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/slider/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
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
	sliders, err := sliderService.sliderRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return sliders, nil
}

func (sliderService sliderServiceImp) GetOneSlider(sliderID int) (*models.Slider, error) {
	slider, err := sliderService.sliderRepo.FindOne(sliderID)
	if err != nil {
		return nil, err
	}
	return slider, nil
}

func (sliderService sliderServiceImp) CreateSlider(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateSliderRequest) error {
	// file upload
	sliderImageURL, err := utils.UploadFile(ctx, "slider_image_url", config.FolderConfig.PublicPath, "slider-images")
	if err != nil {
		return err
	}
	// create model
	createSlider := models.Slider{
		SliderImageURL: sliderImageURL,
		SliderStatus:   "ACTIVE",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// create slider
	if err := sliderService.sliderRepo.Store(createSlider); err != nil {
		return nil
	}

	return nil
}

func (sliderService sliderServiceImp) UpdateSliderStatus(sliderID int, updateSlider dto.UpdateSliderStatus) error {

	slider, err := sliderService.sliderRepo.FindOne(sliderID)
	if err != nil {
		return err
	}
	slider.SliderStatus = updateSlider.SliderStatus
	slider.UpdatedAt = time.Now()

	if err := sliderService.sliderRepo.UpdateStatus(slider.ID, *slider); err != nil {
		return err
	}
	return nil

}

func (sliderService sliderServiceImp) DeleteSlider(sliderID int) error {
	slider, err := sliderService.sliderRepo.FindOne(sliderID)
	if err != nil {
		return err
	}

	if err := sliderService.sliderRepo.Destroy(slider.ID); err != nil {
		return err
	}
	return nil
}
