package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/category/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type CategoryService interface {
	// admin for functions
	GetOneCategoryWithID(categoryID int) (*dto.GetOneCategoryResponse, error)
	GetAllCategory() ([]dto.GetAllCategoryResponse, error)
	CreateCategory(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateCategoryRequest) error
	UpdateCategory(ctx *fiber.Ctx, config *config.Config, categoryID int, updateRequest dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error

	// front for functions

	GetOneCategoryWithSlug(categorySlug string) (*dto.GetOneCategoryResponse, error)
}
