package dto

type CreateSliderRequest struct {
	SliderImageURL string `form:"slider_image_url"`
}

type UpdateSliderStatus struct {
	SliderStatus string `form:"slider_status" validate:"required"`
}
