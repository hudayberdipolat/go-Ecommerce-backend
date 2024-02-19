package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ProductService interface {
	GetOneProduct(productID int) (*dto.OneProductResponse, error)
	GetAllProduct() ([]dto.AllProductResponse, error)
	CreateProduct(ctx *fiber.Ctx, config *config.Config, request dto.CreateProductRequest) error
	UpdateProduct(ctx *fiber.Ctx, config *config.Config, productID int, request dto.UpdateProductRequest) error
	DeleteProduct(productID int) error
}
