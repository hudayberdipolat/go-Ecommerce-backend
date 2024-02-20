package dto

type AboutRequest struct {
	AboutDesc       string `json:"about_desc" validate:"required"`
	LocationMapHtml string `json:"location_map_html" validate:"required"`
}
