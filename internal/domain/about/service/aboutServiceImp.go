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
	about, err := aboutService.aboutRepo.GetOne(aboutID)
	if err != nil {
		return nil, err
	}

	aboutResponse := dto.NewAboutResponse(about)
	return &aboutResponse, nil
}

func (aboutService aboutServiceImp) UpdateAbout(aboutID int, about models.About) error {

	updateAbout, err := aboutService.aboutRepo.GetOne(aboutID)
	if err != nil {
		return err
	}

	// update about data
	updateAbout.AboutDesc = about.AboutDesc
	updateAbout.LocationMapHtml = about.LocationMapHtml

	// update about
	if err := aboutService.aboutRepo.Update(updateAbout.ID, *updateAbout); err != nil {
		return err
	}
	return nil
}
