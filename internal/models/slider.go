package models

import "time"

type Slider struct {
	ID             int       `json:"id"`
	SliderImageURL string    `json:"slider_image_url"`
	SliderStatus   string    `json:"slider_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (*Slider) TableName() string {
	return "sliders"
}
