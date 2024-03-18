package dto

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type GetOneProductResponse struct {
	ID                  int                `json:"id"`
	ProductNameTk       string             `json:"product_name_tk"`
	ProductNameRu       string             `json:"product_name_ru"`
	ProductNameEn       string             `json:"product_name_en"`
	ProductSlug         string             `json:"product_slug"`
	ProductShortDescTk  string             `json:"product_short_desc_tk"`
	ProductShortDescRU  string             `json:"product_short_desc_ru"`
	ProductShortDescEn  string             `json:"product_short_desc_en"`
	ProductMainImageURL *string            `json:"product_main_image_url"`
	ProductLongDescTk   string             `json:"product_long_desc_tk"`
	ProductLongDescRu   string             `json:"product_long_desc_ru"`
	ProductLongDescEn   string             `json:"product_long_desc_en"`
	ProductModel        string             `json:"product_model"`
	ProductStatus       string             `json:"product_status"`
	ProductFeature      string             `json:"product_feature"`
	OriginalPrice       float64            `json:"original_price"`
	DisCountPrice       float64            `json:"discount_price"`
	DisCountTime        string             `json:"discount_time"`
	TotalCount          int                `json:"total_count"`
	RestCount           int                `json:"rest_count"`
	ProductImages       []productImage     `json:"product_images"`
	Category            productCategory    `json:"category"`
	SubCategory         productSubCategory `json:"sub_category"`
	Brand               productBrand       `json:"brand"`
	ProductComment      []productComment   `json:"product_comment"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
}

type productCategory struct {
	ID               int     `json:"id"`
	CategoryNameTk   string  `json:"category_name_tk"`
	CategoryNameRu   string  `json:"category_name_ru"`
	CategoryNameEn   string  `json:"category_name_en"`
	CategoryImageURL *string `json:"category_image_url"`
	CategorySlug     string  `json:"category_slug"`
	CategoryStatus   string  `json:"category_status"`
	CreatedAt        string  `json:"created_at"`
}

type productSubCategory struct {
	ID                  int     `json:"id"`
	SubCategoryNameTk   string  `json:"sub_category_tk"`
	SubCategoryNameRu   string  `json:"sub_category_ru"`
	SubCategoryNameEn   string  `json:"sub_category_en"`
	SubCategorySlug     string  `json:"sub_category_slug"`
	SubCategoryImageURL *string `json:"sub_categorty_image_url"`
	SubCategoryStatus   string  `json:"sub_category_status"`
	CreatedAt           string  `json:"created_at"`
}

type productBrand struct {
	ID            int     `json:"id"`
	BrandNameTk   string  `json:"brand_name_tk"`
	BrandNameRu   string  `json:"brand_name_ru"`
	BrandNameEn   string  `json:"brand_name_en"`
	BrandImageURL *string `json:"brand_image_url"`
	BrandSlug     string  `json:"brand_slug"`
	BrandStatus   string  `json:"brand_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type productComment struct {
	ID       int    `json:"id"`
	Comment  string `json:"comment"`
	UserName string `json:"username"`
}

type productImage struct {
	ID       int    `json:"id"`
	ImageUrl string `json:"image_url"`
}

func NewGetOneProductResponse(product *models.Product) GetOneProductResponse {

	var productImages []productImage
	var productComments []productComment
	for _, image := range product.ProductImages {
		productImage := productImage{
			ID:       image.ID,
			ImageUrl: *image.ImageURL,
		}
		productImages = append(productImages, productImage)
	}

	for _, comment := range product.ProductComments {
		productComment := productComment{
			ID:       comment.ID,
			Comment:  comment.ProductComment,
			UserName: comment.User.Username,
		}
		productComments = append(productComments, productComment)
	}

	return GetOneProductResponse{
		ID:                  product.ID,
		ProductNameTk:       product.ProductNameTk,
		ProductNameRu:       product.ProductNameRu,
		ProductNameEn:       product.ProductNameEn,
		ProductSlug:         product.ProductSlug,
		ProductShortDescTk:  product.ProductShortDescTk,
		ProductShortDescRU:  product.ProductShortDescRu,
		ProductShortDescEn:  product.ProductShortDescEn,
		ProductLongDescTk:   product.ProductLongDescTk,
		ProductLongDescRu:   product.ProductLongDescRu,
		ProductLongDescEn:   product.ProductLongDescEn,
		ProductMainImageURL: product.ProductMainImageURL,
		ProductModel:        product.ProductModel,
		ProductStatus:       product.ProductStatus,
		ProductFeature:      product.ProductFeature,
		OriginalPrice:       product.OriginalPrice,
		DisCountPrice:       product.DisCountPrice,
		TotalCount:          product.TotalCount,
		RestCount:           product.RestCount,
		ProductImages:       productImages,
		ProductComment:      productComments,
		Category: productCategory{
			ID:               product.Category.ID,
			CategoryNameTk:   product.Category.CategoryNameTk,
			CategoryNameRu:   product.Category.CategoryNameRu,
			CategoryNameEn:   product.Category.CategoryNameEn,
			CategorySlug:     product.Category.CategorySlug,
			CategoryImageURL: product.Category.CategoryImageURL,
			CategoryStatus:   product.Category.CategoryStatus,
			CreatedAt:        product.Category.CreatedAt.Format("01-02-2006"),
		},
		SubCategory: productSubCategory{
			ID:                  product.SubCategory.ID,
			SubCategoryNameTk:   product.SubCategory.SubCategoryNameTk,
			SubCategoryNameRu:   product.SubCategory.SubCategoryNameRu,
			SubCategoryNameEn:   product.SubCategory.SubCategoryNameEn,
			SubCategorySlug:     product.SubCategory.SubCategorySlug,
			SubCategoryImageURL: product.SubCategory.SubCategoryImageURL,
			SubCategoryStatus:   product.SubCategory.SubCategoryStatus,
			CreatedAt:           product.SubCategory.Category.CreatedAt.Format("01-02-2006"),
		},
		Brand: productBrand{
			ID:            product.Brand.ID,
			BrandNameTk:   product.Brand.BrandNameTk,
			BrandNameRu:   product.Brand.BrandNameRu,
			BrandNameEn:   product.Brand.BrandNameEn,
			BrandSlug:     product.Brand.BrandSlug,
			BrandImageURL: product.Brand.BrandImageURL,
			BrandStatus:   product.Brand.BrandStatus,
			CreatedAt:     product.Brand.CreatedAt.Format("01-02-2006"),
			UpdatedAt:     product.Brand.UpdatedAt.Format("01-02-2006"),
		},
		CreatedAt: product.CreatedAt.Format("01-02-2006"),
		UpdatedAt: product.UpdatedAt.Format("01-02-2006"),
	}

}

type GetAllProductResponse struct {
	ID                  int     `json:"id"`
	ProductNameTk       string  `json:"product_name_tk"`
	ProductNameRu       string  `json:"product_name_ru"`
	ProductNameEn       string  `json:"product_name_en"`
	ProductSlug         string  `json:"product_slug"`
	ProductMainImageURL *string `json:"product_main_image_url"`
	ProductModel        string  `json:"product_model"`
	ProductStatus       string  `json:"product_status"`
	OriginalPrice       float64 `json:"original_price"`
	TotalCount          int     `json:"total_count"`
	RestCount           int     `json:"rest_count"`
}

func NewGetAllProductResponse(products []models.Product) []GetAllProductResponse {
	var productResponses []GetAllProductResponse
	for _, product := range products {
		productResponse := GetAllProductResponse{
			ID:                  product.ID,
			ProductNameTk:       product.ProductNameTk,
			ProductNameRu:       product.ProductNameRu,
			ProductNameEn:       product.ProductNameEn,
			ProductSlug:         product.ProductSlug,
			ProductMainImageURL: product.ProductMainImageURL,
			ProductModel:        product.ProductModel,
			ProductStatus:       product.ProductStatus,
			OriginalPrice:       product.OriginalPrice,
			TotalCount:          product.TotalCount,
			RestCount:           product.RestCount,
		}
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}
