package models

type ProductImage struct {
	ID        int     `json:"id" gorm:"column:id"`
	ImageURL  *string `json:"image_url" gorm:"column:image_url"`
	ProductID int     `json:"product_id" gorm:"column:product_id"`
	Product   Product `json:"product"`
}

func (*ProductImage) TableName() string {
	return "product_images"
}
