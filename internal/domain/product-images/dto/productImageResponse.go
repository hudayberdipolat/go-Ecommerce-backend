package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type ProductImageResponse struct {
	ID       int     `json:"id"`
	ImageURL *string `json:"image_url"`
}

func NewGetAllProductImages(productImages []models.ProductImage) []ProductImageResponse {
	var productImageResponses []ProductImageResponse
	for _, image := range productImages {
		productImageResponse := ProductImageResponse{
			ID:       image.ID,
			ImageURL: image.ImageURL,
		}
		productImageResponses = append(productImageResponses, productImageResponse)
	}
	return productImageResponses
}
