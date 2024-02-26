package service

import (
	"errors"
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/dto"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/domain/admin/repository"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
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

func (adminService adminServiceImp) GetAllAdmin() ([]dto.GetAllAdminResponse, error) {
	admins, err := adminService.adminRepo.FindAll()
	if err != nil {
		return nil, err
	}
	allAdminResponse := dto.NewGetAllAdminResponse(admins)
	return allAdminResponse, nil
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*dto.GetOneAdminResponse, error) {
	admin, err := adminService.adminRepo.FindOneAdminWithID(adminID)
	if err != nil {
		return nil, err
	}
	adminResponse := dto.NewGetOneAdminResponse(admin)
	return &adminResponse, nil
}

func (adminService adminServiceImp) CreateAdmin(request dto.CreateAdminRequest) error {

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MaxCost)
	createAdmin := models.Admin{
		Firstname:   request.Firstname,
		Surname:     request.Surname,
		PhoneNumber: request.PhoneNumber,
		AdminRole:   "active",
		AdminStatus: "admin",
		Password:    string(passwordHash),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := adminService.adminRepo.Create(createAdmin); err != nil {
		return err
	}

	return nil
}

func (adminService adminServiceImp) UpdateAdminData(adminID int, requet dto.UpdateAdminDataRequest) error {
	updateAdmin, err := adminService.adminRepo.FindOneAdminWithID(adminID)
	if err != nil {
		return err
	}
	// update admin data
	updateAdmin.Firstname = requet.Firstname
	updateAdmin.Surname = requet.Surname
	updateAdmin.PhoneNumber = requet.PhoneNumber
	updateAdmin.AdminRole = requet.AdminRole
	updateAdmin.AdminStatus = requet.AdminStatus
	updateAdmin.UpdatedAt = time.Now()

	if err := adminService.adminRepo.UpdateAdminData(updateAdmin.ID, *updateAdmin); err != nil {
		return err
	}

	return nil

}

func (adminService adminServiceImp) UpdateAdminPassword(adminID int, requet dto.UpdateAdminPasswordRequest) error {
	updateAdminPassword, err := adminService.adminRepo.FindOneAdminWithID(adminID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(updateAdminPassword.Password), []byte(requet.OldPassword)); err != nil {
		return errors.New("Köne password nädogry!!!")
	}

	if err := adminService.adminRepo.UpdateAdminPassword(updateAdminPassword.ID, requet.Password); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	deleteAdmin, err := adminService.adminRepo.FindOneAdminWithID(adminID)
	if err != nil {
		return err
	}

	// delete admin
	if err := adminService.adminRepo.Delete(deleteAdmin.ID); err != nil {
		return err
	}
	return nil
}

func (adminService adminServiceImp) LoginAdmin(request dto.AdminLoginRequest) (*dto.AdminAuthResponse, error) {
	// get admin data with phoneNumber
	getAdmin, err := adminService.adminRepo.FindAdminWithPhoneNumber(request.PhoneNumber)
	if err != nil {
		return nil, errors.New("Telefon belgi ýa-da password nädogry!!!")
	}

	// password barlag
	if err := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("Telefon belgi ýa-da password nädogry!!!")
	}

	// generate admin token
	accessToken, err := adminToken.GenerateAdminToken(getAdmin.ID, getAdmin.PhoneNumber, getAdmin.AdminRole, getAdmin.AdminStatus)

	if err != nil {
		return nil, err
	}

	// return admin response
	adminResponse := dto.NewAdminAuthResponse(getAdmin, accessToken)
	return &adminResponse, nil

}
