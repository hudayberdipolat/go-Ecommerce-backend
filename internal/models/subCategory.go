package models

import "time"

type SubCategory struct {
	ID                  int
	SubCategoryNameTk   string    `json:"sub_category_tk"`
	SubCategoryNameRu   string    `json:"sub_category_ru"`
	SubCategoryNameEn   string    `json:"sub_category_en"`
	SubCategoryImageURL *string   `json:"sub_category_image_url"`
	SubCategorySlug     string    `json:"sub_category_slug"`
	SubCategoryStatus   string    `json:"sub_category_status"`
	CategoryID          int       `json:"category_id"`
	Category            Category  `json:"category" gorm:"foreignKey:CategoryID;references:ID"` // BelongsTo is correct
	Products            []Product `json:"products"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (*SubCategory) TableName() string {
	return "sub_categories"
}
