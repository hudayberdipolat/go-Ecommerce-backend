package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type BrendService interface {
	GetOneBrend(brendID int) (*dto.OneBrendResponse, error)
	GetAllBrend() ([]dto.AllBrendResponse, error)
	CreateBrend(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateBrendRequest) error
	UpdateBrend(ctx *fiber.Ctx, config config.Config, brendID int, updateRequest dto.UpdateBrendRequest) error
	DeleteBrend(brendID int) error
}
