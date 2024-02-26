package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type OneProductResponse struct {
	ID            int                     `json:"id"`
	ProductNameTk string                  `json:"product_name_tk"`
	ProductNameRu string                  `json:"product_name_ru"`
	ProductSlug   string                  `json:"product_slug"`
	ProductDescTk string                  `json:"product_desc_tk"`
	ProductDescRu string                  `json:"product_desc_ru"`
	ProductStatus string                  `json:"product_status"`
	MainImage     *string                 `json:"Main_image"`
	Price         string                  `json:"price"`
	TotalCount    int                     `json:"total_count"`
	GalanSany     int                     `json:"galan_sany"`
	Category      ProductCategoryResponse `json:"category"`
	Brend         ProductBrendResponse    `json:"brend"`
	ProductImages []ProductImagesResponse `json:"product_images"`
}

type ProductImagesResponse struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"product_id"`
	ProductImage string `json:"product_image"`
}

type ProductCategoryResponse struct {
	ID             int    `json:"id"`
	CategoryNameTk string `json:"category_name_tk"`
	CategoryNameRu string `json:"category_name_ru"`
	CategorySlug   string `json:"category_slug"`
}

type ProductBrendResponse struct {
	ID          int    `json:"id"`
	BrendNameTk string `json:"brend_name_tk"`
	BrendNameRu string `json:"brend_name_ru"`
	BrendImage  string `json:"brend_image"`
	BrendSlug   string `json:"brend_slug"`
}

func NewOneProductResponse(product *models.Product) OneProductResponse {

	var imagesResponse []ProductImagesResponse

	for _, image := range product.ProductImages {
		imageResponse := ProductImagesResponse{
			ID:           image.ID,
			ProductID:    image.ProductID,
			ProductImage: image.ProductImage,
		}
		imagesResponse = append(imagesResponse, imageResponse)
	}

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
		Category: ProductCategoryResponse{
			ID:             product.Category.ID,
			CategoryNameTk: product.Category.CategoryNameTK,
			CategoryNameRu: product.Category.CategoryNameRU,
			CategorySlug:   product.Category.CategorySlug,
		},
		Brend: ProductBrendResponse{
			ID:          product.Brend.ID,
			BrendNameTk: product.Brend.BrendNameTk,
			BrendNameRu: product.Brend.BrendNameRu,
			BrendImage:  product.Brend.BrendImage,
			BrendSlug:   product.Brend.BrendSlug,
		},
		ProductImages: imagesResponse,
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
