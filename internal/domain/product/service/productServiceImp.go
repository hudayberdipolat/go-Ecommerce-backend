package service

import (
	"github.com/gofiber/fiber/v2"
	brandRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
	categoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	subCategoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productServiceImp struct {
	productRepo     productRepository.ProductRepository
	categoryRepo    categoryRepository.CategoryRepository
	subCategoryRepo subCategoryRepository.SubCategoryRepository
	brandRepo       brandRepository.BrandRepository
}

func NewProductService(productRepo productRepository.ProductRepository, categoryRepo categoryRepository.CategoryRepository,
	subCategoryRepo subCategoryRepository.SubCategoryRepository, brandRepo brandRepository.BrandRepository) ProductService {
	return productServiceImp{
		productRepo:     productRepo,
		categoryRepo:    categoryRepo,
		subCategoryRepo: subCategoryRepo,
		brandRepo:       brandRepo,
	}
}

func (productService productServiceImp) GetOneProductByID(productID int) (*dto.GetOneProductResponse, error) {
	panic("product Service Imp")
}

func (productService productServiceImp) GetAllProduct() ([]dto.GetAllProductResponse, error) {
	panic("product Service Imp")
}

func (productService productServiceImp) CreateProduct(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateProductRequest) error {
	panic("product Service Imp")
}

func (productService productServiceImp) UpdateProduct(ctx *fiber.Ctx, config *config.Config, productID int, updateRequest dto.UpdateProductRequest) error {
	panic("product Service Imp")
}

func (productService productServiceImp) DeleteProduct(productID int) error {
	panic("product Service Imp")
}
