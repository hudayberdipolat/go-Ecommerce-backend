package models

type Brend struct {
	ID          int    `json:"id"`
	BrendNameTk string `json:"brend_name_tk"`
	BrendNameRu string `json:"brend_name_ru"`
	BrendSlug   string `json:"brend_slug"`
	BrendStatus string `json:"brend_status"`
	BrendImage  string `json:"brend_image"`
}

func (*Brend) TableName() string {
	return "brends"
}
