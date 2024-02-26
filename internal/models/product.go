package models

type Product struct {
	ID            int            `json:"id"`
	ProductNameTk string         `json:"product_name_tk"`
	ProductNameRu string         `json:"product_name_ru"`
	ProductSlug   string         `json:"product_slug"`
	ProductDescTk string         `json:"product_desc_tk"`
	ProductDescRu string         `json:"product_desc_ru"`
	ProductStatus string         `json:"product_status"`
	MainImage     *string        `json:"Main_image"`
	Price         string         `json:"price"`
	TotalCount    int            `json:"total_count"`
	GalanSany     int            `json:"galan_sany"`
	CategoryID    int            `json:"category_id"`
	Category      Category       `json:"category"`
	BrendID       int            `json:"brend_id"`
	Brend         Brend          `json:"brend"`
	ProductImages []ProductImage `json:"images"`
}

func (*Product) TableName() string {
	return "products"
}
