package dto

type CreateAppDataRequest struct {
	AppLogo    *string `form:"app_logo"`
	PlayMarket *string `form:"play_market"`
	AppStore   *string `form:"app_store"`
	QrCode     *string `form:"qr_code"`
}
type UpdateAppDataRequest struct {
	AppLogo    *string `form:"app_logo"`
	PlayMarket *string `form:"play_market"`
	AppStore   *string `form:"app_store"`
	QrCode     *string `form:"qr_code"`
}
