package dto

type CreateBrandRequest struct {
	BrandNameTk   string `form:"brand_name_tk" validate:"required,min=3"`
	BrandNameRu   string `form:"brand_name_ru" validate:"required,min=3"`
	BrandNameEn   string `form:"brand_name_en" validate:"required,min=3"`
	BrandImageURL string `form:"brand_image_url"`
}

type UpdateBrandRequest struct {
	BrandNameTk   string `form:"brand_name_tk" validate:"required,min=3"`
	BrandNameRu   string `form:"brand_name_ru" validate:"required,min=3"`
	BrandNameEn   string `form:"brand_name_en" validate:"required,min=3"`
	BrandImageURL string `form:"brand_image_url" validate:"required"`
	BrandSlug     string `form:"brand_slug" validate:"required"`
	BrandStatus   string `form:"brand_status" validate:"required"`
}
