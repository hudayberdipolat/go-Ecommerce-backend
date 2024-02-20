package repository

import (
	"errors"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/models"
	"gorm.io/gorm"
)

type brendRepositoryImp struct {
	db *gorm.DB
}

func NewBrendRepository(db *gorm.DB) BrendRepository {
	return brendRepositoryImp{db: db}
}

func (b brendRepositoryImp) FindOneBrend(brendID int) (*models.Brend, error) {
	var brend models.Brend
	if err := b.db.First(&brend, brendID).Error; err != nil {
		return nil, err
	}
	return &brend, nil

}

func (b brendRepositoryImp) FindAllBrend() ([]models.Brend, error) {
	var brends []models.Brend
	if err := b.db.Find(&brends).Error; err != nil {
		return nil, err
	}
	return brends, nil
}

func (b brendRepositoryImp) Create(brend models.Brend) error {
	if err := b.db.Create(&brend).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu Brend ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (b brendRepositoryImp) Update(brendID int, brend models.Brend) error {
	err := b.db.Model(models.Brend{}).Where("id=?", brendID).Updates(&brend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu Brend ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (b brendRepositoryImp) Delete(brendID int) error {
	if err := b.db.Unscoped().Delete(&models.Brend{}, brendID).Error; err != nil {
		return err
	}
	return nil
}

func (b brendRepositoryImp) CheckBrendName(brendNameTk, brendNameRu string) bool {
	var brend models.Brend
	b.db.Where("brend_name_tk =? OR brend_name_ru=?", brendNameTk, brendNameRu).First(&brend)
	if brend.ID == 0 {
		return true
	}
	return false
}
