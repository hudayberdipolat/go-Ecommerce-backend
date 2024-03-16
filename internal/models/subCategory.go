package models

import "time"

type SubCategory struct {
	ID                  int       `json:"id"`
	SubCategoryNameTk   string    `json:"sub_category_name_tk" gorm:"column:sub_category_name_tk"`
	SubCategoryNameRu   string    `json:"sub_category_name_ru" gorm:"column:sub_category_name_ru"`
	SubCategoryNameEn   string    `json:"sub_category_name_en" gorm:"column:sub_category_name_en"`
	SubCategoryImageURL *string   `json:"sub_category_image_url" gorm:"column:sub_category_image_url"`
	SubCategorySlug     string    `json:"sub_category_slug" gorm:"column:sub_category_slug"`
	SubCategoryStatus   string    `json:"sub_category_status" gorm:"column:sub_category_status"`
	CategoryID          int       `json:"category_id" gorm:"column:category_id"`
	Category            Category  `json:"category" gorm:"foreignKey:CategoryID;references:ID;category_id"` // BelongsTo is correct
	Products            []Product `json:"products"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ProductCount        int       `json:"product_count" gorm:"->" `
}

func (*SubCategory) TableName() string {
	return "sub_categories"
}
