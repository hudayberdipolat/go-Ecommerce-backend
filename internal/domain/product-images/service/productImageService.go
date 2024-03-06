package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ProductImageService interface {
	GetAllProductImages(productID int) ([]dto.ProductImageResponse, error)
	GetOneProductImage(productID, productImageID int) (*dto.ProductImageResponse, error)
	CreateProductImage(ctx *fiber.Ctx, config config.Config, productID int, request dto.CreateProductRequest) error
	DeleteProductImage(productID, productImageID int) error
}
