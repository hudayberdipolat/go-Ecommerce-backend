package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/productImage/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type ProductImageService interface {
	CreateProductImage(ctx *fiber.Ctx, config *config.Config, request dto.ProductImageRequest) error
	Delete(productID, productImageID int) error
}
