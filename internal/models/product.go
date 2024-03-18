package models

import (
	"time"
)

type Product struct {
	ID                  int              `json:"id" gorm:"column:id"`
	ProductNameTk       string           `json:"product_name_tk" gorm:"column:product_name_tk"`
	ProductNameRu       string           `json:"product_name_ru" gorm:"column:product_name_ru"`
	ProductNameEn       string           `json:"product_name_en" gorm:"column:product_name_en"`
	ProductSlug         string           `json:"product_slug" gorm:"column:product_slug"`
	ProductShortDescTk  string           `json:"product_short_desc_tk" gorm:"column:product_short_desc_tk"`
	ProductShortDescRu  string           `json:"product_short_desc_ru" gorm:"column:product_short_desc_ru"`
	ProductShortDescEn  string           `json:"product_short_desc_en" gorm:"column:product_short_desc_en"`
	ProductLongDescTk   string           `json:"product_long_desc_tk" gorm:"column:product_long_desc_tk"`
	ProductLongDescRu   string           `json:"product_long_desc_ru" gorm:"column:product_long_desc_ru"`
	ProductLongDescEn   string           `json:"product_long_desc_en" gorm:"column:product_long_desc_en"`
	ProductMainImageURL *string          `json:"product_main_image_url" gorm:"column:product_main_image_url"`
	ProductStatus       string           `json:"product_status" gorm:"column:product_status"`
	ProductModel        string           `json:"product_model" gorm:"column:product_model"`
	ProductFeature      string           `json:"product_feature" gorm:"column:product_feature"`
	OriginalPrice       float64          `json:"original_price" gorm:"column:original_price"`
	DisCountPrice       float64          `json:"discount_price" gorm:"column:discount_price"`
	TotalCount          int              `json:"total_count" gorm:"column:total_count"`
	RestCount           int              `json:"rest_count" gorm:"column:rest_count"`
	CategoryID          int              `json:"category_id" gorm:"column:category_id"`
	SubCategoryID       int              `json:"sub_category_id" gorm:"column:sub_category_id"`
	BrandID             int              `json:"brand_id" gorm:"column:brand_id"`
	Category            Category         `json:"category"`
	SubCategory         SubCategory      `json:"sub_category"`
	Brand               Brand            `json:"brand"`
	ProductImages       []ProductImage   `json:"product_images"   gorm:"->"`
	ProductComments     []ProductComment `json:"product_comments" gorm:"->"`
	CreatedAt           time.Time        `json:"created_at"`
	UpdatedAt           time.Time        `json:"updated_at"`
	// CommentCount        int              `json:"comment_count" gorm:"->"`
}

func (*Product) TableName() string {
	return "products"
}
