package dto

type CreateCategoryRequest struct {
	CategoryNameTk   string `form:"category_name_tk" validate:"required,min=3"`
	CategoryNameRu   string `form:"category_name_ru" validate:"required,min=3"`
	CategoryNameEn   string `form:"category_name_en" validate:"required,min=3"`
	CategoryImageURL string `form:"category_image_url"`
}

type UpdateCategoryRequest struct {
	CategoryNameTk   string `form:"category_name_tk" validate:"required,min=3"`
	CategoryNameRu   string `form:"category_name_ru" validate:"required,min=3"`
	CategoryNameEn   string `form:"category_name_en" validate:"required,min=3"`
	CategoryImageURL string `form:"category_image_url"`
	CategoryStatus   string `form:"category_status" validate:"required"`
}
