package service

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/jwtToken/adminToken"
	"golang.org/x/crypto/bcrypt"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (adminService adminServiceImp) GetAllAdmin() ([]models.Admin, error) {
	admins, err := adminService.adminRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*models.Admin, error) {
	getAdmin, err := adminService.adminRepo.FindOne(adminID)
	if err != nil {
		return nil, err
	}
	return getAdmin, nil
}

func (adminService adminServiceImp) CreateAdmin(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateAdminRequest) error {

	// generate password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	image, _ := ctx.FormFile("admin_image_url")

	var imageUrl *string
	if image != nil {
		// upload image
		adminImageURL, err := utils.UploadFile(ctx, "admin_image_url", config.FolderConfig.PublicPath, "admin-images")
		if err != nil {
			return err
		}
		imageUrl = adminImageURL
	}

	createAdmin := models.Admin{
		Username:      createRequest.Username,
		FullName:      createRequest.FullName,
		PhoneNumber:   createRequest.PhoneNumber,
		Email:         createRequest.Email,
		AdminStatus:   "ACTIVE",
		AdminRole:     "admin",
		AdminImageURL: imageUrl,
		Password:      string(hashPassword),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := adminService.adminRepo.Store(createAdmin); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) UpdateAdmin(ctx *fiber.Ctx, config *config.Config, adminID int, updateRequest dto.UpdateAdminRequest) error {
	updateAdmin, err := adminService.adminRepo.FindOne(adminID)
	if err != nil {
		return err
	}

	file, _ := ctx.FormFile("admin_image_url")

	if file != nil {
		// old image delete
		if updateAdmin.AdminImageURL != nil {
			if err := utils.DeleteFile(*updateAdmin.AdminImageURL); err != nil {
				return err
			}
		}
		// new image upload
		imageUrl, err := utils.UploadFile(ctx, "admin_image_url", config.FolderConfig.PublicPath, "admin-images")
		if err != nil {
			return err
		}
		updateAdmin.AdminImageURL = imageUrl
	}

	updateAdmin.Username = updateRequest.Username
	updateAdmin.FullName = updateRequest.FullName
	updateAdmin.Email = updateRequest.Email
	updateAdmin.AdminStatus = updateRequest.AdminStatys

	if err := adminService.adminRepo.Update(updateAdmin.ID, *updateAdmin); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) UpdateAdminPassword(adminID int, updatePassword dto.ChangeAdminPasswordRequest) error {
	admin, err := adminService.adminRepo.FindOne(adminID)
	if err != nil {
		return err
	}
	// admin password denesdirmek

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(updatePassword.OldPassword)); err != nil {
		return errors.New("old password wrong")
	}

	// hash password

	password, _ := bcrypt.GenerateFromPassword([]byte(updatePassword.Password), bcrypt.DefaultCost)

	if err := adminService.adminRepo.UpdateAdminPassword(admin.ID, string(password)); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	admin, err := adminService.adminRepo.FindOne(adminID)
	if err != nil {
		return err
	}
	if err := adminService.adminRepo.Destroy(admin.ID); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) Login(adminLoginRequest dto.LoginAdminRequest) (*dto.AdminAuthResponse, error) {

	getAdmin, err := adminService.adminRepo.FindAdminWithPhoneNumber(adminLoginRequest.PhoneNumber)
	if err != nil {
		return nil, errors.New("phone number or password wrong")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(adminLoginRequest.Password)); err != nil {
		return nil, errors.New("phone number or password wrong")
	}
	accessToken, errToken := adminToken.GenerateAdminToken(getAdmin.ID, getAdmin.PhoneNumber, getAdmin.AdminRole, getAdmin.AdminStatus)
	if errToken != nil {
		return nil, err
	}
	adminResponse := dto.NewAdminAuthResponse(getAdmin, accessToken)
	return &adminResponse, nil
}
