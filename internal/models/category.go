package models

type Category struct {
	ID             int    `json:"id"`
	CategoryNameTK string `json:"category_name_tk"`
	CategoryNameRU string `json:"category_name_ru"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `joson:"category_status"`
}

func (*Category) TableName() string {
	return "categories"
}
