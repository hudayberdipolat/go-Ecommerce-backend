package models

type Product struct {
	ID            int    `json:"id"`
	ProductName   string `json:"product_name"`
	ProductSlug   string `json:"product_slug"`
	ProductDesc   string `json:"product_desc"`
	ImageUrl      string `json:"image_url"`
	Price         string `json:"price"`
	TotalCount    int    `json:"total_count"`
	GalanSany     int    `json:"galan_sany"`
	ProductStatus string `json:"product_status"`
}

func (*Product) TableName() string {
	return "products"
}
