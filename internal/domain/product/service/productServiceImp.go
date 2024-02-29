package service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/product/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type productServiceImp struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return productServiceImp{
		productRepo: repo,
	}
}

func (p productServiceImp) GetOneProduct(productID int) (*dto.OneProductResponse, error) {
	product, err := p.productRepo.FindOne(productID)
	if err != nil {
		return nil, err
	}
	productResponse := dto.NewOneProductResponse(product)
	return &productResponse, nil
}

func (p productServiceImp) GetAllProduct() ([]dto.AllProductResponse, error) {
	products, err := p.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	productResponses := dto.NewAllProductResponse(products)
	return productResponses, nil
}

func (p productServiceImp) CreateProduct(ctx *fiber.Ctx, config *config.Config, request dto.CreateProductRequest) error {

	// file upload
	path, err := utils.UploadFile(ctx, "main_image", config.FolderConfig.PublicPath, "product-images")
	if err != nil {
		return err
	}
	createProduct := models.Product{
		ProductNameTk: request.ProductNameTk,
		ProductNameRu: request.ProductNameRu,
		ProductSlug:   slug.Make(request.ProductNameTk),
		ProductStatus: request.ProductStatus,
		ProductDescTk: request.ProductDescTk,
		ProductDescRu: request.ProductDescRu,
		MainImage:     path,
		Price:         request.Price,
		TotalCount:    request.TotalCount,
		GalanSany:     request.TotalCount,
		CategoryID:    request.CategoryID,
		BrendID:       request.BrendID,
	}

	if err := p.productRepo.Create(createProduct); err != nil {
		return err
	}
	return nil
}

func (p productServiceImp) UpdateProduct(ctx *fiber.Ctx, config *config.Config, productID int, request dto.UpdateProductRequest) error {
	updateProduct, err := p.productRepo.FindOne(productID)
	if err != nil {
		return errors.New("product not found")
	}

	file, _ := ctx.FormFile("main_image")
	if file != nil {
		// old image delete
		if err := utils.DeleteFile(*updateProduct.MainImage); err != nil {
			return err
		}
		// new image upload
		path, errFileUpload := utils.UploadFile(ctx, "main_image", config.FolderConfig.PublicPath, "product-images")
		if errFileUpload != nil {
			return errFileUpload
		}
		updateProduct.MainImage = path
	}

	updateProduct.ProductNameTk = request.ProductNameTk
	updateProduct.ProductNameRu = request.ProductNameRu
	updateProduct.ProductSlug = slug.Make(request.ProductNameTk)
	updateProduct.ProductStatus = request.ProductStatus
	updateProduct.ProductDescTk = request.ProductDescTk
	updateProduct.ProductDescRu = request.ProductDescRu
	updateProduct.MainImage = updateProduct.MainImage
	updateProduct.ProductNameTk = request.Price
	updateProduct.TotalCount = request.TotalCount
	updateProduct.GalanSany = request.GalanSany
	updateProduct.CategoryID = request.CategoryID
	updateProduct.BrendID = request.BrendID

	if err := p.productRepo.Update(updateProduct.ID, *updateProduct); err != nil {
		return err
	}
	return nil
}

func (p productServiceImp) DeleteProduct(productID int) error {
	deleteProduct, err := p.productRepo.FindOne(productID)
	if err != nil {
		return errors.New("product not found")
	}

	// product image delete

	if err := utils.DeleteFile(*deleteProduct.MainImage); err != nil {
		return err
	}

	if err := p.productRepo.Delete(deleteProduct.ID); err != nil {
		return err
	}
	return nil
}

func (p productServiceImp) GetOneProductWithSlug(productSlug string) (*dto.OneProductResponse, error) {
	getProduct, err := p.productRepo.FindOneProductWithSlug(productSlug)
	if err != nil {
		return nil, err
	}

	// product response
	productResponse := dto.NewOneProductResponse(getProduct)
	return &productResponse, nil
}
