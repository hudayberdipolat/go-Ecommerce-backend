package service

import "github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/dto"

type BrendService interface {
	GetOneBrend(brendID int) (*dto.OneBrendResponse, error)
	GetAllBrend() ([]dto.AllBrendResponse, error)
	CreateBrend(createRequest dto.CreateBrendRequest) error
	UpdateBrend(brendID int, updateRequest dto.UpdateBrendRequest) error
	DeleteBrend(brendID int) error
}
