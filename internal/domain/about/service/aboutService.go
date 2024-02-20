package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/dto"
)

type AboutService interface {
	GetOneAbout(aboutID int) (*dto.AboutResponse, error)
	UpdateAbout(aboutID int, about dto.AboutRequest) error
}
