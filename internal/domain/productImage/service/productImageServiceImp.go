package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/repository"
)

type productImageServiceImp struct {
	productImageRepo repository.ProductImageRepository
}

func NewProductImageService(repo repository.ProductImageRepository) ProductImageService {
	return productImageServiceImp{
		productImageRepo: repo,
	}
}

func (productImageSerive productImageServiceImp) CreateProductImage(request dto.ProductImageRequest) error {
	panic("product image service")
}

func (productImageSerive productImageServiceImp) Delete(productImageID int) error {
	panic("product image service")
}
