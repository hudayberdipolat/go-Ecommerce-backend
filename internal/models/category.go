package models

import (
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/validate"
)

type Category struct {
	ID               int           `json:"id" gorm:"primary_key"`
	CategoryNameTk   string        `json:"category_name_tk"`
	CategoryNameRu   string        `json:"category_name_ru"`
	CategoryNameEn   string        `json:"category_name_en"`
	CategoryImageURL *string       `json:"category_image_url"`
	CategorySlug     string        `json:"category_slug"`
	CategoryStatus   string        `json:"category_status"`
	SubCategories    []SubCategory `json:"sub_categories" gorm:"foreignKey:CategoryID;->"`
	Products         []Product     `json:"products"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	ProductCount     int           `json:"product_count" gorm:"-"`
	SubCategoryCount int           `json:"sub_category_count" gorm:"-"`
}

func (*Category) TableName() string {
	return "categories"
}

func ValidateCatagory(category *Category) error {
	err := validate.ValidateStruct(category)
	return err
}
