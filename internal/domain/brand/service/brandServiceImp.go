package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
)

type brandServiceImp struct {
	brandRepo repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) BrandService {
	return brandServiceImp{
		brandRepo: repo,
	}
}

func (brandService brandServiceImp) GetOneBrandByID(brandID int) (*dto.GetOneBrandResponse, error) {
	panic("brand service imp ")
}

func (brandService brandServiceImp) GetAllBrand() ([]dto.GetAllBrandResponse, error) {
	panic("brand service imp ")
}

func (brandService brandServiceImp) CreateBrand(createRequest dto.CreateBrandRequest) error {
	panic("brand service imp ")
}

func (brandService brandServiceImp) UpdateBrand(brandID int, updateRequest dto.UpdateBrandRequest) error {
	panic("brand service imp ")
}

func (brandService brandServiceImp) DeleteBrand(brandID int) error {
	panic("brand service imp ")
}
