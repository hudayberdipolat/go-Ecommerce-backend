package dto

import "time"

type CreateProductRequest struct {
	ProductNameTk       string    `form:"product_name_tk" validate:"required"`
	ProductNameRu       string    `form:"product_name_ru" validate:"required"`
	ProductNameEn       string    `form:"product_name_en" validate:"required"`
	ProductShortDescTk  string    `form:"product_short_desc_tk" validate:"required"`
	ProductShortDescRu  string    `form:"product_short_desc_ru" validate:"required"`
	ProductShortDescEn  string    `form:"product_short_desc_en" validate:"required"`
	ProductLongDescTk   string    `form:"product_long_desc_tk" validate:"required"`
	ProductLongDescRu   string    `form:"product_long_desc_ru" validate:"required"`
	ProductLongDescEn   string    `form:"product_long_desc_en" validate:"required"`
	ProductMainImageURL string    `form:"product_main_image_url"`
	ProductModel        string    `form:"product_model" validate:"required"`
	OriginalPrice       float32   `form:"original_price" validate:"required"`
	DisCountPrice       float32   `form:"discount_price" validate:"required"`
	DisCountTime        time.Time `form:"discount_time" validate:"required"`
	TotalCount          int       `form:"total_count" validate:"required"`
	CategoryID          int       `form:"category_id" validate:"required"`
	SubCategoryID       int       `form:"sub_category_id" validate:"required"`
	BrandID             int       `form:"brand_id" validate:"required"`
}

type UpdateProductRequest struct {
	ProductNameTk       string    `form:"product_name_tk" validate:"required"`
	ProductNameRu       string    `form:"product_name_ru" validate:"required"`
	ProductNameEn       string    `form:"product_name_en" validate:"required"`
	ProductShortDescTk  string    `form:"product_short_desc_tk" validate:"required"`
	ProductShortDescRu  string    `form:"product_short_desc_ru" validate:"required"`
	ProductShortDescEn  string    `form:"product_short_desc_en" validate:"required"`
	ProductLongDescTk   string    `form:"product_long_desc_tk" validate:"required"`
	ProductLongDescRu   string    `form:"product_long_desc_ru" validate:"required"`
	ProductLongDescEn   string    `form:"product_long_desc_en" validate:"required"`
	ProductMainImageURL string    `form:"product_main_image_url"`
	ProductStatus       string    `form:"product_status" validate:"required"`
	ProductModel        string    `form:"product_model" validate:"required"`
	OriginalPrice       float32   `form:"original_price" validate:"required"`
	DisCountPrice       float32   `form:"discount_price" validate:"required"`
	DisCountTime        time.Time `form:"discount_time" validate:"required"`
	TotalCount          int       `form:"total_count" validate:"required"`
	RestCount           int       `form:"rest_count" validate:"required"`
	CategoryID          int       `form:"category_id" validate:"required"`
	SubCategoryID       int       `form:"sub_category_id" validate:"required"`
	BrandID             int       `form:"brand_id" validate:"required"`
}
