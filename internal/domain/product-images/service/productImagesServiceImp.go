package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/dto"
	pImageRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product-images/repository"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productImageServiceImp struct {
	pImageRepo  pImageRepository.ProductImageRepository
	productRepo productRepository.ProductRepository
}

func NewProductImageService(repo pImageRepository.ProductImageRepository, pRepo productRepository.ProductRepository) ProductImageService {
	return productImageServiceImp{
		pImageRepo:  repo,
		productRepo: pRepo,
	}
}

func (pImageService productImageServiceImp) GetAllProductImages() ([]dto.ProductImageResponse, error) {
	panic("pImage Service imp")
}

func (pImageService productImageServiceImp) GetOneProductImage(productID, productImageID int) (dto.ProductImageResponse, error) {
	panic("pImage Service imp")
}

func (pImageService productImageServiceImp) CreateProductImage(ctx *fiber.Ctx, config config.Config, productID int, request dto.CreateProductRequest) error {
	panic("pImage Service imp")
}

func (pImageService productImageServiceImp) DeleteProductImage(productID, productImageID int) error {
	panic("pImage Service imp")
}
