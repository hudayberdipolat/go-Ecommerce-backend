package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createRequest.Password), bcrypt.MaxCost)
	if err != nil {
		return err
	}

	// upload image

	adminImageURL, err := utils.UploadFile(ctx, "admin_image_url", config.FolderConfig.PublicPath, "admin-images")
	if err != nil {
		return err
	}

	createAdmin := models.Admin{
		Username:      createRequest.Username,
		FullName:      createRequest.FullName,
		PhoneNumber:   createRequest.PhoneNumber,
		Email:         createRequest.Email,
		AdminStatus:   "ACTIVE",
		AdminImageURL: adminImageURL,
		Password:      string(hashPassword),
		CrearedAt:     time.Now(),
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
		if err := utils.DeleteFile(*updateAdmin.AdminImageURL); err != nil {
			return err
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

	if err := adminService.adminRepo.UpdateAdminPassword(admin.ID, updatePassword.Password); err != nil {
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
