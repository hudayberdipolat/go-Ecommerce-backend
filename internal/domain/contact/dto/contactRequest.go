package dto

type UpdateContactRequest struct {
	PhoneNumber      string `form:"phone_number"`
	YouTubeAccount   string `form:"youtube_account"`
	InstagramAccount string `form:"instagram_account"`
	TiktokAccount    string `form:"tiktok_account"`
	ImoAccount       string `form:"imo_account"`
	Address          string `form:"address"`
	PlayMarket       string `form:"play_market"`
	AppStore         string `form:"app_store"`
	QrCode           string `form:"qr_code"`
}
