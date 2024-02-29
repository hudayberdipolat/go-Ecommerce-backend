package service

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/brend/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
)

type brendServiceImp struct {
	brendRepo repository.BrendRepository
}

func NewBrendService(repo repository.BrendRepository) BrendService {
	return brendServiceImp{brendRepo: repo}
}

func (b brendServiceImp) GetOneBrend(brendID int) (*dto.OneBrendResponse, error) {
	brend, err := b.brendRepo.FindOneBrend(brendID)
	if err != nil {
		return nil, err
	}
	brendResponse := dto.NewOneBrendResponse(brend)
	return &brendResponse, nil
}

func (b brendServiceImp) GetAllBrend() ([]dto.AllBrendResponse, error) {
	brends, err := b.brendRepo.FindAllBrend()
	if err != nil {
		return nil, err
	}

	allBrendResponse := dto.NewAllBrendResponse(brends)
	return allBrendResponse, nil
}

func (b brendServiceImp) CreateBrend(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateBrendRequest) error {

	// check brend Name
	checkBrendName := b.brendRepo.CheckBrendName(createRequest.BrendNameTk, createRequest.BrendNameRu)

	// egerde brend ady on bar bolsa onda onda return error
	if !checkBrendName {
		log.Println(checkBrendName)
		return errors.New("Bu brend ady eýýäm ulanylýar!!!")
	}
	// brend ady on yok bolsa onda image upload image
	// image upload

	path, err := utils.UploadFile(ctx, "brend_image", config.FolderConfig.PublicPath, "brend-images")
	if err != nil {
		return err
	}
	// create brend
	if createRequest.BrendStatus == "" {
		createRequest.BrendStatus = "draft"
	}
	createBrend := models.Brend{
		BrendNameTk: createRequest.BrendNameTk,
		BrendNameRu: createRequest.BrendNameRu,
		BrendSlug:   slug.Make(createRequest.BrendNameTk),
		BrendStatus: createRequest.BrendStatus,
		BrendImage:  *path,
	}

	if err := b.brendRepo.Create(createBrend); err != nil {
		return err
	}

	return nil
}

func (b brendServiceImp) UpdateBrend(ctx *fiber.Ctx, config config.Config, brendID int, updateRequest dto.UpdateBrendRequest) error {
	updateBrend, err := b.brendRepo.FindOneBrend(brendID)
	if err != nil {
		return errors.New("brend not found")
	}

	file, _ := ctx.FormFile("brend_image")
	if file != nil {
		// old image delete
		if err := utils.DeleteFile(updateBrend.BrendImage); err != nil {
			return err
		}
		// new image upload
		path, errFileUpload := utils.UploadFile(ctx, "brend_image", config.FolderConfig.PublicPath, "brend-images")
		if errFileUpload != nil {
			return errFileUpload
		}
		updateRequest.BrendImage = *path
	}

	updateBrend.BrendNameTk = updateRequest.BrendNameTk
	updateBrend.BrendNameRu = updateRequest.BrendNameRu
	updateBrend.BrendSlug = slug.Make(updateRequest.BrendNameTk)
	updateBrend.BrendStatus = updateRequest.BrendStatus
	updateBrend.BrendImage = updateRequest.BrendImage
	if err := b.brendRepo.Update(updateBrend.ID, *updateBrend); err != nil {
		return err
	}
	return nil
}

func (b brendServiceImp) DeleteBrend(brendID int) error {
	deleteBrend, err := b.brendRepo.FindOneBrend(brendID)
	if err != nil {
		return errors.New("Brend Not Found")
	}

	// image delete
	if err := utils.DeleteFile(deleteBrend.BrendImage); err != nil {
		return err
	}
	// delete brend

	if err := b.brendRepo.Delete(deleteBrend.ID); err != nil {
		return err
	}
	return nil
}

func (b brendServiceImp) GetOneBrendWithSlug(brendSlug string) (*dto.OneBrendResponse, error) {
	brend, err := b.brendRepo.GetOneBrendWithSlug(brendSlug)
	if err != nil {
		return nil, err
	}
	brendResponse := dto.NewOneBrendResponse(brend)
	return &brendResponse, nil
}
