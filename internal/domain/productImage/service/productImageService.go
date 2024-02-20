package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/dto"

type ProductImageService interface {
	CreateProductImage(request dto.ProductImageRequest) error
	Delete(productImageID int) error
}
