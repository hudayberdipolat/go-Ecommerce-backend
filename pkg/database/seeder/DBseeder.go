package seeder

import (
	"time"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"github.com/pkg/errors"
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

	password, errPass := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if errPass != nil {
		return errors.Wrap(errPass, "failed to generate password hash")
	}

	// admins seeder
	admins := []models.Admin{
		{
			Username:      "polat",
			FullName:      "Hudayberdi Polatov",
			PhoneNumber:   "99365097512",
			Email:         "hudayberdipolat@gmail.com",
			AdminStatus:   "ACTIVE",
			AdminRole:     "super_admin",
			AdminImageURL: nil,
			Password:      string(password),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Username:      "hudayberdi",
			FullName:      "Polat Polatov",
			PhoneNumber:   "99365010203",
			Email:         "hudayberdipolat@gmail.com",
			AdminStatus:   "ACTIVE",
			AdminRole:     "admin",
			AdminImageURL: nil,
			Password:      string(password),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	for _, admin := range admins {
		if err := s.db.Create(&admin).Error; err != nil {
			return errors.Wrap(err, "failed to seed admins")
		}
	}

	// user seeder
	users := []models.User{
		{
			Username:    "user",
			PhoneNumber: "99365010203",
			Email:       "user@gmail.com",
			UserStatus:  "ACTIVE",
			Password:    string(password),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Username:    "user2",
			PhoneNumber: "99365030201",
			Email:       "user2@gmail.com",
			UserStatus:  "ACTIVE",
			Password:    string(password),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Username:    "user3",
			PhoneNumber: "99365020103",
			Email:       "user3@gmail.com",
			UserStatus:  "ACTIVE",
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
