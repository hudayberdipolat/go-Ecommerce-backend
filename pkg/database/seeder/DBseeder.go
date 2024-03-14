package seeder

import (
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NewSeederInterface interface {
	Seeder() error
}

type newSeederImp struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) NewSeederInterface {
	return newSeederImp{
		db: db,
	}
}

func (s newSeederImp) Seeder() error {

	password, errPass := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MaxCost)
	if errPass != nil {
		return errPass
	}

	// admins seeder
	admins := []models.Admin{
		{
			Username:    "polat",
			FullName:    "Hudayberdi Polatov",
			PhoneNumber: "99365097512",
			Email:       "hudayberdipolat@gmail.com",
			// AdminImageURL: nil,
			AdminStatus: "active",
			Password:    string(password),
			CrearedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := s.db.Create(&admins).Error; err != nil {
		return nil
	}

	// user seeder

	users := []models.User{
		{
			Username:    "user",
			PhoneNumber: "99365010203",
			Email:       "user@gmail.com",
			UserStatus:  "active",
			Password:    string(password),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Username:    "user2",
			PhoneNumber: "99365030201",
			Email:       "user2@gmail.com",
			UserStatus:  "active",
			Password:    string(password),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Username:    "user3",
			PhoneNumber: "99365020103",
			Email:       "user3@gmail.com",
			UserStatus:  "active",
			Password:    string(password),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := s.db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}
