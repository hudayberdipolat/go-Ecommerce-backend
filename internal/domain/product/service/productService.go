package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ProductService interface {
	GetOneProduct(productID int) (*dto.OneProductResponse, error)
	GetAllProduct() ([]dto.AllProductResponse, error)
	CreateProduct(ctx *fiber.Ctx, config config.Config, request models.Product) error
	UpdateProduct(ctx *fiber.Ctx, config config.Config, productID int, request models.Product) error
	DeleteProduct(productID int) error
}
