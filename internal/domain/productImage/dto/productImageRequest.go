package dto

type ProductImageRequest struct {
	ProductID    int    `json:"product_id" form:"product_id"`
	ProductImage string `json:"product_image" form:"product_image"`
}
