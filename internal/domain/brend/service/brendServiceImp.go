package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/repository"
)

type brendServiceImp struct {
	brendRepo repository.BrendRepository
}

func NewBrendService(repo repository.BrendRepository) BrendService {
	return brendServiceImp{brendRepo: repo}
}

func (b brendServiceImp) GetOneBrend(brendID int) (*dto.OneBrendResponse, error) {
	panic("brend service imp")
}

func (b brendServiceImp) GetAllBrend() ([]dto.AllBrendResponse, error) {
	panic("brend service imp")
}

func (b brendServiceImp) CreateBrend(createRequest dto.CreateBrendRequest) error {
	panic("brend service imp")
}

func (b brendServiceImp) UpdateBrend(brendID int, updateRequest dto.UpdateBrendRequest) error {
	panic("brend service imp")
}

func (b brendServiceImp) DeleteBrend(brendID int) error {
	panic("brend service imp")
}
