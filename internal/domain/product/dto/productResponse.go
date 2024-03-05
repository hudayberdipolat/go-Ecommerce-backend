package dto

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type GetOneProductResponse struct {
	ProductNameTk       string             `json:"product_name_tk"`
	ProductNameRu       string             `json:"product_name_ru"`
	ProductNameEn       string             `json:"product_name_en"`
	ProductSlug         string             `json:"product_slug"`
	ProductShortDescTk  string             `json:"product_short_desc_tk"`
	ProductShortDescRU  string             `json:"product_short_desc_ru"`
	ProductShortDescEn  string             `json:"product_short_desc_en"`
	ProductMainImageURL *string            `json:"product_main_image_url"`
	ProductModel        string             `json:"product_model"`
	ProductStatus       string             `json:"product_status"`
	OriginalPrice       float32            `json:"original_price"`
	DisCountPrice       float32            `json:"discount_price"`
	DisCountTime        string             `json:"discount_time"`
	TotalCount          int                `json:"total_count"`
	RestCount           int                `json:"rest_count"`
	Category            productCategory    `json:"category"`
	SubCategory         productSubCategory `json:"sub_category"`
	Brand               productBrand       `json:"brand"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
}

type productCategory struct {
	ID               int     `json:"id"`
	CategoryNameTk   string  `json:"category_name_tk"`
	CategoryNameRu   string  `json:"category_name_ru"`
	CategoryNameEn   string  `json:"category_name_en"`
	CategoryImageURL *string `json:"category_image_url"`
	CategorySlug     string  `json:"category_slug"`
	CategoryStatus   string  `json:"category_status"`
	CreatedAt        string  `json:"created_at"`
}

type productSubCategory struct {
	ID                  int     `json:"id"`
	SubCategoryNameTk   string  `json:"sub_category_tk"`
	SubCategoryNameRu   string  `json:"sub_category_ru"`
	SubCategoryNameEn   string  `json:"sub_category_en"`
	SubCategorySlug     string  `json:"sub_category_slug"`
	SubCategoryImageURL *string `json:"sub_categorty_image_url"`
	SubCategoryStatus   string  `json:"sub_category_status"`
	CreatedAt           string  `json:"created_at"`
}

type productBrand struct {
	ID            int     `json:"id"`
	BrandNameTk   string  `json:"brand_name_tk"`
	BrandNameRu   string  `json:"brand_name_ru"`
	BrandNameEn   string  `json:"brand_name_en"`
	BrandImageURL *string `json:"brand_image_url"`
	BrandSlug     string  `json:"brand_slug"`
	BrandStatus   string  `json:"brand_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func NewGetOneProductResponse(product *models.Product) GetOneProductResponse {
	return GetOneProductResponse{
		ProductNameTk:       product.ProductNameTk,
		ProductNameRu:       product.ProductNameRu,
		ProductNameEn:       product.ProductNameEn,
		ProductSlug:         product.ProductSlug,
		ProductShortDescTk:  product.ProductShortDescTk,
		ProductShortDescRU:  product.ProductShortDescRu,
		ProductShortDescEn:  product.ProductShortDescEn,
		ProductMainImageURL: product.ProductMainImageURL,
		ProductModel:        product.ProductModel,
		ProductStatus:       product.ProductStatus,
		OriginalPrice:       product.OriginalPrice,
		DisCountPrice:       product.DisCountPrice,
		DisCountTime:        product.DisCountTime.Format("01-02-2006"),
		TotalCount:          product.TotalCount,
		RestCount:           product.RestCount,
		Category: productCategory{
			ID:               product.Category.ID,
			CategoryNameTk:   product.Category.CategoryNameTk,
			CategoryNameRu:   product.Category.CategoryNameRu,
			CategoryNameEn:   product.Category.CategoryNameEn,
			CategorySlug:     product.Category.CategorySlug,
			CategoryImageURL: product.Category.CategoryImageURL,
			CategoryStatus:   product.Category.CategoryStatus,
			CreatedAt:        product.Category.CreatedAt.Format("01-02-2006"),
		},
		SubCategory: productSubCategory{
			ID:                  product.SubCategory.ID,
			SubCategoryNameTk:   product.SubCategory.SubCategoryNameTk,
			SubCategoryNameRu:   product.SubCategory.SubCategoryNameRu,
			SubCategoryNameEn:   product.SubCategory.SubCategoryNameEn,
			SubCategorySlug:     product.SubCategory.SubCategorySlug,
			SubCategoryImageURL: product.SubCategory.SubCategoryImageURL,
			SubCategoryStatus:   product.SubCategory.SubCategoryStatus,
			CreatedAt:           product.SubCategory.Category.CreatedAt.Format("01-02-2006"),
		},
		Brand: productBrand{
			ID:            product.Brand.ID,
			BrandNameTk:   product.Brand.BrandNameTk,
			BrandNameRu:   product.Brand.BrandNameRu,
			BrandNameEn:   product.Brand.BrandNameEn,
			BrandSlug:     product.Brand.BrandSlug,
			BrandImageURL: product.Brand.BrandImageURL,
			BrandStatus:   product.Brand.BrandStatus,
			CreatedAt:     product.Brand.CreatedAt.Format("01-02-2006"),
			UpdatedAt:     product.Brand.UpdatedAt.Format("01-02-2006"),
		},
		CreatedAt: product.CreatedAt.Format("01-02-2006"),
		UpdatedAt: product.UpdatedAt.Format("01-02-2006"),
	}
}

type GetAllProductResponse struct {
	ProductNameTk       string  `json:"product_name_tk"`
	ProductNameRu       string  `json:"product_name_ru"`
	ProductNameEn       string  `json:"product_name_en"`
	ProductSlug         string  `json:"product_slug"`
	ProductMainImageURL *string `json:"product_main_image_url"`
	ProductModel        string  `json:"product_model"`
	ProductStatus       string  `json:"product_status"`
	OriginalPrice       float32 `json:"original_price"`
	TotalCount          int     `json:"total_count"`
	RestCount           int     `json:"rest_count"`
}

func NewGetAllProductResponse(products []models.Product) []GetAllProductResponse {
	var productResponses []GetAllProductResponse
	for _, product := range products {
		productResponse := GetAllProductResponse{
			ProductNameTk:       product.ProductNameTk,
			ProductNameRu:       product.ProductNameRu,
			ProductNameEn:       product.ProductNameEn,
			ProductSlug:         product.ProductSlug,
			ProductMainImageURL: product.ProductMainImageURL,
			ProductModel:        product.ProductModel,
			ProductStatus:       product.ProductStatus,
			OriginalPrice:       product.OriginalPrice,
			TotalCount:          product.TotalCount,
			RestCount:           product.RestCount,
		}
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}
