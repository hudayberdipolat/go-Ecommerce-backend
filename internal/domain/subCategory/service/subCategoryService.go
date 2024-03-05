package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/subCategory/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type SubCategoryService interface {
	GetOneSubCategory(subCategoryID int) (*dto.GetOneSubCategoryResponse, error)
	GetAllSubCategory() ([]dto.GetAllSubCategoryResponse, error)
	CreateSubCategory(ctx *fiber.Ctx, config config.Config, categoryID int, createSubCategory dto.CreateSubCategoryRequest) error
	UpdateSubCategory(ctx *fiber.Ctx, config config.Config, categoryID int, subCategoryID int, updateSubCategory dto.UpdateSubCategoryRequest) error
	DeleteSubCategory(categoryID, subCategoryID int) error
}
