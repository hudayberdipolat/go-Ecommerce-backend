package models

type ProductImage struct {
	ID           int     `json:"id"`
	ProductID    int     `json:"product_id"`
	ProductImage string  `json:"product_image"`
	Product      Product `json:"product"`
}

func (*ProductImage) TableName() string {
	return "product_images"
}
