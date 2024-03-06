package models

type ProductImage struct {
	ID        int     `json:"id"`
	ImageURL  *string `json:"image_url"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
}

func (*ProductImage) TableName() string {
	return "product_images"
}
