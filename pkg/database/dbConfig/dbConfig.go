package dbconfig

import (
	"fmt"

	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	config *config.Config
}

func NewDBConnection(config *config.Config) *DBConfig {
	return &DBConfig{
		config: config,
	}
}

func (dbConfig DBConfig) GetDBConfig() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.config.DBConfig.DbHost,
		dbConfig.config.DBConfig.DbUser,
		dbConfig.config.DBConfig.DbPassword,
		dbConfig.config.DBConfig.DbName,
		dbConfig.config.DBConfig.DbPort,
		dbConfig.config.DBConfig.DbSllMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}
	return db, err
}
