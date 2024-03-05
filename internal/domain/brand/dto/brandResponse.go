package dto

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type GetOneBrandResponse struct {
	ID            int                    `json:"id"`
	BrandNameTk   string                 `json:"brand_name_tk"`
	BrandNameRu   string                 `json:"brand_name_ru"`
	BrandNameEn   string                 `json:"brand_name_en"`
	BrandImageURL string                 `json:"brand_image_url"`
	BrandSlug     string                 `json:"brand_slug"`
	BrandStatus   string                 `json:"brand_status"`
	Products      []brandProductResponse `json:"products"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

type brandProductResponse struct {
	ID              int     `json:"id"`
	ProductNameTk   string  `json:"product_name_tk"`
	ProductNameRu   string  `json:"product_name_ru"`
	ProductNameEn   string  `json:"product_name_en"`
	ProductSlug     string  `json:"product_slug"`
	ProductImageURL string  `json:"product_image_url"`
	ProductStatus   string  `json:"product_status"`
	OriginalPrice   float32 `json:"original_price"`
	CreatedAt       string  `json:"created_at"`
}

func NewGetOneBrandResponse(brand models.Brand) GetOneBrandResponse {
	var productResponses []brandProductResponse

	for _, product := range brand.Products {
		brandProduct := brandProductResponse{
			ID:              product.ID,
			ProductNameTk:   product.ProductNameTk,
			ProductNameRu:   product.ProductNameRu,
			ProductNameEn:   product.ProductNameEn,
			ProductSlug:     product.ProductSlug,
			ProductImageURL: product.ProductMainImageURL,
			ProductStatus:   product.ProductStatus,
			OriginalPrice:   product.OriginalPrice,
			CreatedAt:       product.CreatedAt.Format("01-02-2006"),
		}
		productResponses = append(productResponses, brandProduct)
	}

	return GetOneBrandResponse{
		ID:            brand.ID,
		BrandNameTk:   brand.BrandNameTk,
		BrandNameRu:   brand.BrandNameRu,
		BrandNameEn:   brand.BrandNameEn,
		BrandSlug:     brand.BrandSlug,
		BrandImageURL: brand.BrandImageURL,
		BrandStatus:   brand.BrandStatus,
		CreatedAt:     brand.CreatedAt.Format("01-02-2006"),
		UpdatedAt:     brand.UpdatedAt.Format("01-02-2006"),
		Products:      productResponses,
	}
}

type GetAllBrandResponse struct {
	ID            int    `json:"id"`
	BrandNameTk   string `json:"brand_name_tk"`
	BrandNameRu   string `json:"brand_name_ru"`
	BrandNameEn   string `json:"brand_name_en"`
	BrandImageURL string `json:"brand_image_url"`
	BrandSlug     string `json:"brand_slug"`
	BrandStatus   string `json:"brand_status"`
	ProductCount  int    `json:"product_count"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func NewGetAllBrandResponse(brands []models.Brand) []GetAllBrandResponse {
	var brandResponses []GetAllBrandResponse
	productCount := 0
	for _, brand := range brands {
		for i := 0; i < len(brand.Products); i++ {
			productCount = productCount + 1
		}
		brandResponse := GetAllBrandResponse{
			ID:            brand.ID,
			BrandNameTk:   brand.BrandNameTk,
			BrandNameRu:   brand.BrandNameRu,
			BrandNameEn:   brand.BrandNameEn,
			BrandSlug:     brand.BrandSlug,
			BrandImageURL: brand.BrandImageURL,
			BrandStatus:   brand.BrandStatus,
			ProductCount:  productCount,
			CreatedAt:     brand.CreatedAt.Format("01-02-2006"),
			UpdatedAt:     brand.UpdatedAt.Format("01-02-2006"),
		}
		brandResponses = append(brandResponses, brandResponse)
	}
	return brandResponses
}
