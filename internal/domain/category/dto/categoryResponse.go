package dto

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type GetOneCategoryResponse struct {
	ID               int                       `json:"id"`
	CategoryNameTk   string                    `json:"category_name_tk"`
	CategoryNameRu   string                    `json:"category_name_ru"`
	CategoryNameEn   string                    `json:"category_name_en"`
	CategoryImageURL *string                   `json:"category_image_url"`
	CategorySlug     string                    `json:"category_slug"`
	CategoryStatus   string                    `json:"category_status"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        string                    `json:"updated_at"`
	SubCategories    []subCategoryResponse     `json:"sub_categories"`
	Products         []categoryProductResponse `json:"products"`
}

type subCategoryResponse struct {
	ID                  int     `json:"id"`
	SubCategoryNameTk   string  `json:"sub_category_tk"`
	SubCategoryNameRu   string  `json:"sub_category_ru"`
	SubCategoryNameEn   string  `json:"sub_category_en"`
	SubCategorySlug     string  `json:"sub_category_slug"`
	SubCategoryImageURL *string `json:"sub_categorty_image_url"`
	SubCategoryStatus   string  `json:"sub_category_status"`
	CreatedAt           string  `json:"created_at"`
}

type categoryProductResponse struct {
	ID              int     `json:"id"`
	ProductNameTk   string  `json:"product_name_tk"`
	ProductNameRu   string  `json:"product_name_ru"`
	ProductNameEn   string  `json:"product_name_en"`
	ProductSlug     string  `json:"product_slug"`
	ProductImageURL *string `json:"product_image_url"`
	ProductStatus   string  `json:"product_status"`
	OriginalPrice   float32 `json:"original_price"`
	CreatedAt       string  `json:"created_at"`
}

func NewGetOneCategoryResponse(category *models.Category) GetOneCategoryResponse {
	var subCategoriesResponses []subCategoryResponse
	var productResponses []categoryProductResponse

	for _, subCategory := range category.SubCategories {
		subCategoryResponse := subCategoryResponse{
			ID:                  subCategory.ID,
			SubCategoryNameTk:   subCategory.SubCategoryNameTk,
			SubCategoryNameRu:   subCategory.SubCategoryNameRu,
			SubCategoryNameEn:   subCategory.SubCategoryNameEn,
			SubCategorySlug:     subCategory.SubCategorySlug,
			SubCategoryImageURL: subCategory.SubCategoryImageURL,
			SubCategoryStatus:   subCategory.SubCategoryStatus,
			CreatedAt:           subCategory.Category.CreatedAt.Format("01-02-2006"),
		}
		subCategoriesResponses = append(subCategoriesResponses, subCategoryResponse)
	}

	for _, product := range category.Products {
		productResponse := categoryProductResponse{
			ID:              product.ID,
			ProductNameTk:   product.ProductNameTk,
			ProductNameRu:   product.ProductNameRu,
			ProductNameEn:   product.ProductNameEn,
			ProductSlug:     product.ProductSlug,
			ProductImageURL: product.ProductMainImageURL,
			ProductStatus:   product.ProductSlug,
			OriginalPrice:   product.OriginalPrice,
			CreatedAt:       product.CreatedAt.Format("01-02-2006"),
		}
		productResponses = append(productResponses, productResponse)
	}

	return GetOneCategoryResponse{
		ID:               category.ID,
		CategoryNameTk:   category.CategoryNameTk,
		CategoryNameRu:   category.CategoryNameRu,
		CategoryNameEn:   category.CategoryNameEn,
		CategoryImageURL: category.CategoryImageURL,
		CategorySlug:     category.CategorySlug,
		CategoryStatus:   category.CategoryStatus,
		CreatedAt:        category.CreatedAt.Format("01-02-2006"),
		UpdatedAt:        category.UpdatedAt.Format("01-02-2006"),
		SubCategories:    subCategoriesResponses,
		Products:         productResponses,
	}
}

type GetAllCategoryResponse struct {
	ID               int     `json:"id"`
	CategoryNameTk   string  `json:"category_name_tk"`
	CategoryNameRu   string  `json:"category_name_ru"`
	CategoryNameEn   string  `json:"category_name_en"`
	CategoryImageURL *string `json:"category_image_url"`
	CategorySlug     string  `json:"category_slug"`
	CategoryStatus   string  `json:"category_status"`
	CreatedAt        string  `json:"created_at"`
}

func NewGetAllCategoryResponse(categories []models.Category) []GetAllCategoryResponse {
	var getAllResponses []GetAllCategoryResponse

	for _, category := range categories {
		categoryResponse := GetAllCategoryResponse{
			ID:               category.ID,
			CategoryNameTk:   category.CategoryNameTk,
			CategoryNameRu:   category.CategoryNameRu,
			CategoryNameEn:   category.CategoryNameEn,
			CategoryImageURL: category.CategoryImageURL,
			CategorySlug:     category.CategorySlug,
			CategoryStatus:   category.CategoryStatus,
			CreatedAt:        category.CreatedAt.Format("01-02-2006"),
		}
		getAllResponses = append(getAllResponses, categoryResponse)
	}
	return getAllResponses
}
