package models

import "time"

type Product struct {
	ID                  int              `json:"id"`
	ProductNameTk       string           `json:"product_name_tk"`
	ProductNameRu       string           `json:"product_name_ru"`
	ProductNameEn       string           `json:"product_name_en"`
	ProductSlug         string           `json:"product_slug"`
	ProductShortDescTk  string           `json:"product_short_desc_tk"`
	ProductShortDescRu  string           `json:"product_short_desc_ru"`
	ProductShortDescEn  string           `json:"product_short_desc_en"`
	ProductLongDescTk   string           `json:"product_long_desc_tk"`
	ProductLongDescRu   string           `json:"product_long_desc_ru"`
	ProductLongDescEn   string           `json:"product_long_desc_en"`
	ProductMainImageURL *string          `json:"product_main_image_url"`
	ProductModel        string           `json:"product_model"`
	ProductStatus       string           `json:"product_status"`
	ProductFeature      string           `json:"product_feature"`
	OriginalPrice       float32          `json:"original_price"`
	DisCountPrice       float32          `json:"discount_price"`
	DisCountTime        time.Time        `json:"discount_time"`
	TotalCount          int              `json:"total_count"`
	RestCount           int              `json:"rest_count"`
	CategoryID          int              `json:"category_id"`
	SubCategoryID       int              `json:"sub_category_id"`
	BrandID             int              `json:"brand_id"`
	Category            Category         `json:"category"`
	SubCategory         SubCategory      `json:"sub_category"`
	Brand               Brand            `json:"brand"`
	ProductImages       []ProductImage   `json:"product_images"`
	ProductComments     []ProductComment `json:"product_comments"`
	CreatedAt           time.Time        `json:"created_at"`
	UpdatedAt           time.Time        `json:"updated_at"`
}
