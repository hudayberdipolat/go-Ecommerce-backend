package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ProductService interface {
	GetOneProductByID(productID int) (*dto.GetOneProductResponse, error)
	GetAllProduct() ([]dto.GetAllProductResponse, error)
	CreateProduct(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateProductRequest) error
	UpdateProduct(ctx *fiber.Ctx, config *config.Config, productID int, updateRequest dto.UpdateProductRequest) error
	DeleteProduct(productID int) error
}
