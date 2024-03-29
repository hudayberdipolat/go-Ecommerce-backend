package service

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	brandRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brand/repository"
	categoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	productRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	subCategoryRepository "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
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
	product, err := productService.productRepo.GetOneByID(productID)
	if err != nil {
		return nil, err
	}

	productResponse := dto.NewGetOneProductResponse(product)
	return &productResponse, nil
}

func (productService productServiceImp) GetAllProduct() ([]dto.GetAllProductResponse, error) {
	products, err := productService.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	productResponses := dto.NewGetAllProductResponse(products)
	return productResponses, nil
}

func (productService productServiceImp) CreateProduct(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateProductRequest) error {
	category, err := productService.categoryRepo.FindOneByID(createRequest.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}
	subCategory, err := productService.subCategoryRepo.FindOne(category.ID, createRequest.SubCategoryID)
	if err != nil {
		return errors.New("subCategory not found")
	}
	brand, err := productService.brandRepo.GetOneByID(createRequest.BrandID)
	if err != nil {
		return errors.New("brand not found")
	}

	productImageURL, err := utils.UploadFile(ctx, "product_main_image_url", config.FolderConfig.PublicPath, "product-images")
	if err != nil {
		return err
	}
	randString := utils.RandStringRunes(6)
	createProduct := models.Product{
		ProductNameTk:       createRequest.ProductNameTk,
		ProductNameRu:       createRequest.ProductNameRu,
		ProductNameEn:       createRequest.ProductNameEn,
		ProductSlug:         slug.Make(createRequest.ProductNameEn) + "-" + randString,
		ProductShortDescTk:  createRequest.ProductShortDescTk,
		ProductShortDescRu:  createRequest.ProductShortDescRu,
		ProductShortDescEn:  createRequest.ProductShortDescEn,
		ProductLongDescTk:   createRequest.ProductLongDescTk,
		ProductLongDescRu:   createRequest.ProductLongDescRu,
		ProductLongDescEn:   createRequest.ProductLongDescEn,
		ProductMainImageURL: productImageURL,
		ProductModel:        createRequest.ProductModel,
		ProductStatus:       "DRAFT",
		OriginalPrice:       createRequest.OriginalPrice,
		DisCountPrice:       createRequest.DisCountPrice,
		TotalCount:          createRequest.TotalCount,
		RestCount:           createRequest.TotalCount,
		CategoryID:          category.ID,
		SubCategoryID:       subCategory.ID,
		BrandID:             brand.ID,
	}

	if err := productService.productRepo.Store(createProduct); err != nil {
		if err := utils.DeleteFile(*productImageURL); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (productService productServiceImp) UpdateProduct(ctx *fiber.Ctx, config *config.Config, productID int, updateRequest dto.UpdateProductRequest) error {

	updateProduct, err := productService.productRepo.GetOneByID(productID)
	if err != nil {
		return err
	}
	category, err := productService.categoryRepo.FindOneByID(updateRequest.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}
	// get subCategory
	subCategory, err := productService.subCategoryRepo.FindOne(category.ID, updateRequest.SubCategoryID)
	if err != nil {
		return errors.New("subCategory not found")
	}
	// get Brend
	brand, err := productService.brandRepo.GetOneByID(updateRequest.BrandID)
	if err != nil {
		return errors.New("brand not found")
	}

	file, _ := ctx.FormFile("product_main_image_url")
	if file != nil {
		// delete file
		if err := utils.DeleteFile(*updateProduct.ProductMainImageURL); err != nil {
			return err
		}
		// file upload
		productImageURL, err := utils.UploadFile(ctx, "product_main_image_url", config.FolderConfig.PublicPath, "product-images")
		if err != nil {
			return err
		}
		updateProduct.ProductMainImageURL = productImageURL
	}
	randString := utils.RandStringRunes(6)

	updateProduct.ProductNameTk = updateRequest.ProductNameTk
	updateProduct.ProductNameRu = updateRequest.ProductNameRu
	updateProduct.ProductNameEn = updateRequest.ProductNameEn
	updateProduct.ProductSlug = slug.Make(updateRequest.ProductNameEn) + "-" + randString
	updateProduct.ProductShortDescTk = updateRequest.ProductShortDescTk
	updateProduct.ProductShortDescTk = updateRequest.ProductShortDescRu
	updateProduct.ProductShortDescTk = updateRequest.ProductShortDescEn
	updateProduct.ProductLongDescTk = updateRequest.ProductLongDescTk
	updateProduct.ProductLongDescTk = updateRequest.ProductLongDescRu
	updateProduct.ProductLongDescTk = updateRequest.ProductLongDescEn
	updateProduct.ProductModel = updateRequest.ProductModel
	updateProduct.ProductStatus = updateRequest.ProductStatus
	updateProduct.OriginalPrice = updateRequest.OriginalPrice
	updateProduct.DisCountPrice = updateRequest.DisCountPrice
	updateProduct.TotalCount = updateRequest.TotalCount
	updateProduct.RestCount = updateRequest.RestCount
	updateProduct.CategoryID = category.ID
	updateProduct.SubCategoryID = subCategory.ID
	updateProduct.BrandID = brand.ID
	updateProduct.UpdatedAt = time.Now()

	if err := productService.productRepo.Update(updateProduct.ID, *updateProduct); err != nil {
		return err
	}
	return nil
}

func (productService productServiceImp) DeleteProduct(productID int) error {
	deleteProduct, err := productService.productRepo.GetOneByID(productID)
	if err != nil {
		return err
	}
	if err := productService.productRepo.Destroy(deleteProduct.ID); err != nil {
		return err
	}
	return nil
}

func (productService productServiceImp) GetOneProductBySlug(productSlug string) (*dto.GetOneProductResponse, error) {
	product, err := productService.productRepo.GetOneBySlug(productSlug)
	if err != nil {
		return nil, err
	}

	productResponse := dto.NewGetOneProductResponse(product)
	return &productResponse, nil
}
