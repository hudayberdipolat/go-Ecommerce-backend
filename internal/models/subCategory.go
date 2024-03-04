package models

type SubCategory struct {
	ID                  int
	SubCategoryTk       string   `json:"sub_category_tk"`
	SubCategoryRu       string   `json:"sub_category_ru"`
	SubCategoryEn       string   `json:"sub_category_en"`
	SubCategoryImageURL string   `json:"sub_category_image_url"`
	SubCategorySlug     string   `json:"sub_category_slug"`
	SubCategoryStatus   string   `json:"sub_category_status"`
	CategoryID          int      `json:"category_id"`
	Category            Category `json:"category"`
}

func (*SubCategory) TableName() string {
	return "sub_categories"
}
