package dto

type CreateCategoryRequest struct {
	CategoryNameTK string `json:"category_name_tk" validate:"required,min=3"`
	CategoryNameRU string `json:"category_name_ru" validate:"required,min=3"`
	CategoryStatus string `json:"category_status" validate:"required"`
}

type UpdateCategoryRequest struct {
	CategoryNameTK string `json:"category_name_tk" validate:"required,min=3"`
	CategoryNameRU string `json:"category_name_ru" validate:"required,min=3"`
	CategoryStatus string `json:"category_status" validate:"required"`
}
