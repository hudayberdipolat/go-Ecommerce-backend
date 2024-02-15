package app

import (
	"log"
	"net/http"

	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	dbconfig "github.com/hudayberdipolat/go-Ecommerce-backend/pkg/database/dbConfig"
	httpCustom "github.com/hudayberdipolat/go-Ecommerce-backend/pkg/http"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/logging"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpServer *http.Client
	Logger     *zerolog.Logger
}

func GetDependencies() (*Dependencies, error) {
	log.Println("logging....")

	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// get logger
	logger := logging.GetLogger(getConfig)
	// db connection
	logger.Info().Msg("-------------------------------------------")
	dbConfig := dbconfig.NewDBConnection(getConfig)
	db, errDB := dbConfig.GetDBConfig()
	if errDB != nil {
		logger.Error().Msg("Database connection error")
		logger.Error().Msg(err.Error())
		return nil, errDB
	}
	logger.Info().Msg("Database connection successfully")
	// http connection
	customHttp := httpCustom.NewHttpClient()
	return &Dependencies{
		Config:     getConfig,
		DB:         db,
		HttpServer: customHttp,
		Logger:     logger,
	}, nil
}
