package models

type About struct {
	ID              int    `json:"id"`
	AboutDesc       string `json:"about_desc"`
	LocationMapHtml string `json:"location_map_html"`
}

func (*About) TableName() string {
	return "about"
}
