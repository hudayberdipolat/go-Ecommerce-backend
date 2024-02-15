package app

import (
	"net/http"

	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	dbconfig "github.com/hudayberdipolat/go-Ecommerce-backend/pkg/database/dbConfig"
	httpCustom "github.com/hudayberdipolat/go-Ecommerce-backend/pkg/http"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpServer *http.Client
}

func GetDependencies() (*Dependencies, error) {

	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	// db connection
	dbConfig := dbconfig.NewDBConnection(getConfig)
	db, errDB := dbConfig.GetDBConfig()
	if errDB != nil {
		return nil, errDB
	}

	// http connection
	customHttp := httpCustom.NewHttpClient()

	return &Dependencies{
		Config:     getConfig,
		DB:         db,
		HttpServer: customHttp,
	}, nil
}
