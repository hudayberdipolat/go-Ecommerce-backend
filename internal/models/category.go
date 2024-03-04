package models

import "time"

type Category struct {
	ID               int       `json:"id"`
	CategoryNameTk   string    `json:"category_name_tk"`
	CategoryNameRu   string    `json:"category_name_ru"`
	CategoryNameEn   string    `json:"category_name_en"`
	CategoryImageURL string    `json:"category_image_url"`
	CategorySlug     string    `json:"category_slug"`
	CategoryStatus   string    `json:"category_status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (*Category) TableName() string {
	return "categories"
}
