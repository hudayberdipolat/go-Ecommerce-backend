package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type AboutService interface {
	GetOneAbout(aboutID int) (*dto.AboutResponse, error)
	UpdateAbout(aboutID int, about models.About) error
}
