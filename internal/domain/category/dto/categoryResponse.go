package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

// admin for response

type OneCategoryResponse struct {
	ID             int
	CategoryNameTK string
	CategoryNameRu string
	CategoryStatus string
	CategorySlug   string
}

func NewOneCategoryResponse(category models.Category) *OneCategoryResponse {
	return &OneCategoryResponse{
		ID:             category.ID,
		CategoryNameTK: category.CategoryNameTK,
		CategoryNameRu: category.CategoryNameRU,
		CategoryStatus: category.CategoryStatus,
		CategorySlug:   category.CategorySlug,
	}
}

type AllCategoryResponse struct{}

func NewAllCategoryResponse(categories []models.Category) []AllCategoryResponse {
	var allCategories []AllCategoryResponse
	return allCategories
}

// user for response
