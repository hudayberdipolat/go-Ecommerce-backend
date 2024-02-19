package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

// admin for response

type OneCategoryResponse struct {
	ID             int    `json:"id"`
	CategoryNameTK string `json:"category_name_tk"`
	CategoryNameRu string `json:"category_name_ru"`
	CategoryStatus string `json:"category_status"`
	CategorySlug   string `json:"category_slug"`
}

func NewOneCategoryResponse(category *models.Category) OneCategoryResponse {
	return OneCategoryResponse{
		ID:             category.ID,
		CategoryNameTK: category.CategoryNameTK,
		CategoryNameRu: category.CategoryNameRU,
		CategoryStatus: category.CategoryStatus,
		CategorySlug:   category.CategorySlug,
	}
}

type AllCategoryResponse struct {
	ID             int    `json:"id"`
	CategoryNameTK string `json:"category_name_tk"`
	CategoryNameRu string `json:"category_name_ru"`
	CategoryStatus string `json:"category_status"`
	CategorySlug   string `json:"category_slug"`
}

func NewAllCategoryResponse(categories []models.Category) []AllCategoryResponse {
	var allCategories []AllCategoryResponse
	for _, category := range categories {
		category := AllCategoryResponse{
			ID:             category.ID,
			CategoryNameTK: category.CategoryNameTK,
			CategoryNameRu: category.CategoryNameRU,
			CategoryStatus: category.CategoryStatus,
			CategorySlug:   category.CategorySlug,
		}
		allCategories = append(allCategories, category)
	}
	return allCategories
}

// user for response
