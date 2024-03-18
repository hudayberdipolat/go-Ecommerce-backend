package models

type AppData struct {
	ID         int     `json:"id"`
	AppLogo    *string `json:"app_logo" gorm:"column:app_logo"`
	PlayMarket *string `json:"play_market" gorm:"column:play_market"`
	AppStore   *string `json:"app_store" gorm:"column:app_store"`
	QrCode     *string `json:"qr_code" gorm:"column:qr_code"`
}

func (AppData) TableName() string {
	return "app_data"
}
