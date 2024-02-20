package service

import (
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/about/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
)

type aboutServiceImp struct {
	aboutRepo repository.AboutRepository
}

func NewAboutService(repo repository.AboutRepository) AboutService {
	return aboutServiceImp{
		aboutRepo: repo,
	}
}

func (aboutService aboutServiceImp) GetOneAbout(aboutID int) (*dto.AboutResponse, error) {
	panic("about service imp")
}

func (aboutService aboutServiceImp) UpdateAbout(aboutID int, about models.About) error {
	panic("about service imp")
}
