package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type BrandService interface {
	GetOneBrandByID(brandID int) (*dto.GetOneBrandResponse, error)
	GetAllBrand() ([]dto.GetAllBrandResponse, error)
	CreateBrand(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateBrandRequest) error
	UpdateBrand(ctx *fiber.Ctx, config *config.Config, brandID int, updateRequest dto.UpdateBrandRequest) error
	DeleteBrand(brandID int) error
}
