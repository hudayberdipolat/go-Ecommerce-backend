package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type GetOneSubCategoryResponse struct {
	ID                  int
	SubCategoryNameTk   string                        `json:"sub_category_tk"`
	SubCategoryNameRu   string                        `json:"sub_category_ru"`
	SubCategoryNameEn   string                        `json:"sub_category_en"`
	SubCategoryImageURL *string                       `json:"sub_category_image_url"`
	SubCategorySlug     string                        `json:"sub_category_slug"`
	SubCategoryStatus   string                        `json:"sub_category_status"`
	Category            categoryResponse              `json:"category" gorm:"foreignKey:CategoryID;references:ID"` // BelongsTo is correct
	Products            []subCategoryProductsResponse `json:"products"`
	CreatedAt           string                        `json:"created_at"`
	UpdatedAt           string                        `json:"updated_at"`
}

type categoryResponse struct {
	ID               int     `json:"id"`
	CategoryNameTk   string  `json:"category_name_tk"`
	CategoryNameRu   string  `json:"category_name_ru"`
	CategoryNameEn   string  `json:"category_name_en"`
	CategoryImageURL *string `json:"category_image_url"`
	CategorySlug     string  `json:"category_slug"`
	CategoryStatus   string  `json:"category_status"`
	CreatedAt        string  `json:"created_at"`
}

type subCategoryProductsResponse struct {
	ID              int     `json:"id"`
	ProductNameTk   string  `json:"product_name_tk"`
	ProductNameRu   string  `json:"product_name_ru"`
	ProductNameEn   string  `json:"product_name_en"`
	ProductSlug     string  `json:"product_slug"`
	ProductImageURL *string `json:"product_image_url"`
	ProductStatus   string  `json:"product_status"`
	OriginalPrice   float64 `json:"original_price"`
	CreatedAt       string  `json:"created_at"`
}

func NewGetOneSubCategoryResponse(subCategory *models.SubCategory) GetOneSubCategoryResponse {
	var productsResponse []subCategoryProductsResponse
	for _, product := range subCategory.Products {
		productResponse := subCategoryProductsResponse{
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
		productsResponse = append(productsResponse, productResponse)

	}
	return GetOneSubCategoryResponse{
		ID:                  subCategory.ID,
		SubCategoryNameTk:   subCategory.SubCategoryNameTk,
		SubCategoryNameRu:   subCategory.SubCategoryNameRu,
		SubCategoryNameEn:   subCategory.SubCategoryNameEn,
		SubCategorySlug:     subCategory.SubCategorySlug,
		SubCategoryImageURL: subCategory.SubCategoryImageURL,
		SubCategoryStatus:   subCategory.SubCategoryStatus,
		Category: categoryResponse{
			ID:               subCategory.Category.ID,
			CategoryNameTk:   subCategory.Category.CategoryNameTk,
			CategoryNameRu:   subCategory.Category.CategoryNameRu,
			CategoryNameEn:   subCategory.Category.CategoryNameEn,
			CategorySlug:     subCategory.Category.CategorySlug,
			CategoryImageURL: subCategory.Category.CategoryImageURL,
			CategoryStatus:   subCategory.Category.CategoryStatus,
			CreatedAt:        subCategory.Category.CreatedAt.Format("01-02-2006"),
		},
		Products:  productsResponse,
		CreatedAt: subCategory.CreatedAt.Format("01-02-2006"),
		UpdatedAt: subCategory.UpdatedAt.Format("01-02-2006"),
	}
}

type GetAllSubCategoryResponse struct {
	ID                  int
	SubCategoryNameTk   string           `json:"sub_category_tk"`
	SubCategoryNameRu   string           `json:"sub_category_ru"`
	SubCategoryNameEn   string           `json:"sub_category_en"`
	SubCategoryImageURL *string          `json:"sub_category_image_url"`
	SubCategorySlug     string           `json:"sub_category_slug"`
	SubCategoryStatus   string           `json:"sub_category_status"`
	Category            categoryResponse `json:"category"`
	ProductCount        int              `json:"product_count"`
	CreatedAt           string           `json:"created_at"`
	UpdatedAt           string           `json:"updated_at"`
}

func NewGetAllSubCategoryResponse(subCategories []models.SubCategory) []GetAllSubCategoryResponse {
	var subCategoryResponses []GetAllSubCategoryResponse
	productCount := 0
	for _, subCategory := range subCategories {
		for i := 0; i < len(subCategory.Products); i++ {
			productCount = productCount + 1
		}

		subCategoryResponse := GetAllSubCategoryResponse{
			ID:                  subCategory.ID,
			SubCategoryNameTk:   subCategory.SubCategoryNameTk,
			SubCategoryNameRu:   subCategory.SubCategoryNameRu,
			SubCategoryNameEn:   subCategory.SubCategoryNameEn,
			SubCategorySlug:     subCategory.SubCategorySlug,
			SubCategoryImageURL: subCategory.SubCategoryImageURL,
			SubCategoryStatus:   subCategory.SubCategoryStatus,
			Category: categoryResponse{
				ID:               subCategory.Category.ID,
				CategoryNameTk:   subCategory.Category.CategoryNameTk,
				CategoryNameRu:   subCategory.Category.CategoryNameRu,
				CategoryNameEn:   subCategory.Category.CategoryNameEn,
				CategorySlug:     subCategory.Category.CategorySlug,
				CategoryImageURL: subCategory.Category.CategoryImageURL,
				CategoryStatus:   subCategory.Category.CategoryStatus,
				CreatedAt:        subCategory.Category.CreatedAt.Format("01-02-2006"),
			},
			ProductCount: productCount,
			CreatedAt:    subCategory.CreatedAt.Format("01-02-2006"),
			UpdatedAt:    subCategory.UpdatedAt.Format("01-02-2006"),
		}
		subCategoryResponses = append(subCategoryResponses, subCategoryResponse)
	}
	return subCategoryResponses
}
