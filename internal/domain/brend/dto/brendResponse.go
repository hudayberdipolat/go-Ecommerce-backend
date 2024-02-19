package dto

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"

type OneBrendResponse struct {
	ID          int    `json:"id"`
	BrendNameTk string `json:"brend_name_tk"`
	BrendNameRu string `json:"brend_name_ru"`
	BrendSlug   string `json:"brend_slug"`
	BrendStatus string `json:"brend_status"`
	BrendImage  string `json:"brend_image"`
}

func NewOneBrendResponse(brend models.Brend) *OneBrendResponse {
	return &OneBrendResponse{
		ID:          brend.ID,
		BrendNameTk: brend.BrendNameTk,
		BrendNameRu: brend.BrendNameRu,
		BrendSlug:   brend.BrendSlug,
		BrendStatus: brend.BrendStatus,
		BrendImage:  brend.BrendImage,
	}
}

type AllBrendResponse struct {
	ID          int    `json:"id"`
	BrendNameTk string `json:"brend_name_tk"`
	BrendNameRu string `json:"brend_name_ru"`
	BrendSlug   string `json:"brend_slug"`
	BrendStatus string `json:"brend_status"`
	BrendImage  string `json:"brend_image"`
}

func NewAllBrendResponse(brends []models.Brend) []AllBrendResponse {
	var allBrends []AllBrendResponse
	for _, brend := range brends {
		brend := AllBrendResponse{
			ID:          brend.ID,
			BrendNameTk: brend.BrendNameTk,
			BrendNameRu: brend.BrendNameRu,
			BrendSlug:   brend.BrendSlug,
			BrendStatus: brend.BrendStatus,
			BrendImage:  brend.BrendImage,
		}
		allBrends = append(allBrends, brend)
	}
	return allBrends
}
