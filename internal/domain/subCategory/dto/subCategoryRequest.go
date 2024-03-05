package dto

type CreateSubCategoryRequest struct {
	SubCategoryNameTk   string `form:"sub_category_name_tk" validate:"required,min=3"`
	SubCategoryNameRu   string `form:"sub_category_name_ru" validate:"required,min=3"`
	SubCategoryNameEn   string `form:"sub_category_name_en" validate:"required,min=3"`
	SubCategoryImageURL string `form:"sub_category_image_url"`
}
type UpdateSubCategoryRequest struct {
	SubCategoryNameTk   string `form:"sub_category_name_tk" validate:"required,min=3"`
	SubCategoryNameRu   string `form:"sub_category_name_ru" validate:"required,min=3"`
	SubCategoryNameEn   string `form:"sub_category_name_en" validate:"required,min=3"`
	SubCategoryImageURL string `form:"sub_category_image_url"`
	SubCategoryStatus   string `form:"sub_category_status" validate:"required"`
}
