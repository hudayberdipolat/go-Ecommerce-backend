package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type OneProductResponse struct {
	ID            int     `json:"id"`
	ProductNameTk string  `json:"product_name_tk"`
	ProductNameRu string  `json:"product_name_ru"`
	ProductSlug   string  `json:"product_slug"`
	ProductDescTk string  `json:"product_desc_tk"`
	ProductDescRu string  `json:"product_desc_ru"`
	ProductStatus string  `json:"product_status"`
	MainImage     *string `json:"Main_image"`
	Price         string  `json:"price"`
	TotalCount    int     `json:"total_count"`
	GalanSany     int     `json:"galan_sany"`
}

func NewOneProductResponse(product *models.Product) OneProductResponse {
	return OneProductResponse{
		ID:            product.ID,
		ProductNameTk: product.ProductNameTk,
		ProductNameRu: product.ProductNameRu,
		ProductSlug:   product.ProductSlug,
		ProductDescTk: product.ProductDescTk,
		ProductDescRu: product.ProductDescRu,
		ProductStatus: product.ProductStatus,
		MainImage:     product.MainImage,
		Price:         product.Price,
		TotalCount:    product.TotalCount,
		GalanSany:     product.GalanSany,
	}
}

type AllProductResponse struct {
	ID            int     `json:"id"`
	ProductNameTk string  `json:"product_name_tk"`
	ProductNameRu string  `json:"product_name_ru"`
	ProductSlug   string  `json:"product_slug"`
	ProductDescTk string  `json:"product_desc_tk"`
	ProductDescRu string  `json:"product_desc_ru"`
	ProductStatus string  `json:"product_status"`
	MainImage     *string `json:"Main_image"`
	Price         string  `json:"price"`
	TotalCount    int     `json:"total_count"`
	GalanSany     int     `json:"galan_sany"`
}

func NewAllProductResponse(products []models.Product) []AllProductResponse {
	var allProductResponse []AllProductResponse

	for _, product := range products {
		productResponse := AllProductResponse{
			ID:            product.ID,
			ProductNameTk: product.ProductNameTk,
			ProductNameRu: product.ProductNameRu,
			ProductSlug:   product.ProductSlug,
			ProductDescTk: product.ProductDescTk,
			ProductDescRu: product.ProductDescRu,
			ProductStatus: product.ProductStatus,
			MainImage:     product.MainImage,
			Price:         product.Price,
			TotalCount:    product.TotalCount,
			GalanSany:     product.GalanSany,
		}
		allProductResponse = append(allProductResponse, productResponse)
	}
	return allProductResponse
}
