package dto

type CreateProductRequest struct {
	ProductNameTk string  `json:"product_name_tk" form:"product_name_tk" validate:"required"`
	ProductNameRu string  `json:"product_name_ru" form:"product_name_ru" validate:"required"`
	ProductDescTk string  `json:"product_desc_tk" form:"product_desc_tk" validate:"required"`
	ProductDescRu string  `json:"product_desc_ru" form:"product_desc_ru" validate:"required"`
	ProductStatus string  `json:"product_status" form:"product_status" validate:"required"`
	MainImage     *string `json:"main_image" form:"main_image"`
	Price         string  `json:"price" form:"price" validate:"required"`
	TotalCount    int     `json:"total_count" form:"total_count" validate:"required"`
	CategoryID    int     `json:"category_id" form:"category_id" validate:"required"`
	BrendID       int     `json:"brend_id" form:"brend_id" validate:"required"`
}

type UpdateProductRequest struct {
	ProductNameTk string  `json:"product_name_tk" form:"product_name_tk" validate:"required"`
	ProductNameRu string  `json:"product_name_ru" form:"product_name_ru" validate:"required"`
	ProductDescTk string  `json:"product_desc_tk" form:"product_desc_tk" validate:"required"`
	ProductDescRu string  `json:"product_desc_ru" form:"product_desc_ru" validate:"required"`
	ProductStatus string  `json:"product_status" form:"product_status" validate:"required"`
	MainImage     *string `json:"main_image" form:"main_image" `
	Price         string  `json:"price" form:"price" validate:"required"`
	TotalCount    int     `json:"total_count" form:"total_count" validate:"required"`
	GalanSany     int     `json:"galan_sany" form:"galan_sany" validate:"required"`
	CategoryID    int     `json:"category_id" form:"category_id" validate:"required"`
	BrendID       int     `json:"brend_id" form:"brend_id" validate:"required"`
}
