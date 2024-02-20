package dto

type CreateBrendRequest struct {
	BrendNameTk string `json:"brend_name_tk" form:"brend_name_tk" validate:"required,min=3"`
	BrendNameRu string `json:"brend_name_ru" form:"brend_name_ru" validate:"required,min=3"`
	BrendStatus string `json:"brend_status" form:"brend_status" validate:"required"`
	BrendImage  string `json:"brend_image" form:"brend_image"`
}

type UpdateBrendRequest struct {
	BrendNameTk string `json:"brend_name_tk" form:"brend_name_tk" validate:"required,min=3"`
	BrendNameRu string `json:"brend_name_ru" form:"brend_name_ru" validate:"required,min=3"`
	BrendStatus string `json:"brend_status" form:"brend_status" validate:"required,min=3"`
	BrendImage  string `json:"brend_image"  form:"brend_image"`
}
