package models

import "time"

type Brand struct {
	ID            int       `json:"id"`
	BrandNameTk   string    `json:"brand_name_tk"`
	BrandNameRu   string    `json:"brand_name_ru"`
	BrandNameEn   string    `json:"brand_name_en"`
	BrandImageURL *string   `json:"brand_image_url"`
	BrandSlug     string    `json:"brand_slug"`
	BrandStatus   string    `json:"brand_status"`
	Products      []Product `json:"products"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ProductCount  int       `json:"product_count" gorm:"-"`
}

func (*Brand) TableName() string {
	return "brands"
}
