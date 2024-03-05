package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/dto"

type BrandService interface {
	GetOneBrandByID(brandID int) (*dto.GetOneBrandResponse, error)
	GetAllBrand() ([]dto.GetAllBrandResponse, error)
	CreateBrand(createRequest dto.CreateBrandRequest) error
	UpdateBrand(brandID int, updateRequest dto.UpdateBrandRequest) error
	DeleteBrand(brandID int) error
}
