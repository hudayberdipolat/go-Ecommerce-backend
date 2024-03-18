package dto

type UpdateContactRequest struct {
	PhoneNumber      string `json:"phone_number" validate:"required"`
	YouTubeAccount   string `json:"youtube_account" validate:"required"`
	InstagramAccount string `json:"instagram_account" validate:"required"`
	TiktokAccount    string `json:"tiktok_account" validate:"required"`
	ImoAccount       string `json:"imo_account" validate:"required"`
	Address          string `json:"address" validate:"required"`
}
