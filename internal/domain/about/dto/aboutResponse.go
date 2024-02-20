package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type AboutResponse struct {
	ID              int    `json:"id"`
	AboutDesc       string `json:"about_desc"`
	LocationMapHtml string `json:"location_map_html"`
}

func NewAboutResponse(about models.About) AboutResponse {
	return AboutResponse{
		ID:              about.ID,
		AboutDesc:       about.AboutDesc,
		LocationMapHtml: about.LocationMapHtml,
	}
}
