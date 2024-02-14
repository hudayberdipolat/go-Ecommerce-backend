package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBConfig     DBConfig     `json:"db_config"`
	HttpConfig   HttpConfig   `json:"http_config"`
	FolderConfig FolderConfig `json:"folder_config"`
}

type DBConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST"`
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSllMode  string `json:"db_sll_mode" env:"DB_SLL_MODE"`
}

type HttpConfig struct {
	ServerHost string `json:"" env:"SERVER_HOST"`
	ServerPort string `json:"" env:"SERVER_PORT"`
	AppName    string `json:"" env:"APP_NAME"`
	AppHeader  string `json:"" env:"APP_HEADER"`
}

type FolderConfig struct {
	RootPath   string `json:"root_path" env:"ROOT_PATH"`
	PublicPath string `json:"public_path" env:"ROOT_PATH"`
}

func GetConfig() (*Config, error) {

	var cfg Config

	if err := cleanenv.ReadConfig("../.env", &cfg); err != nil {
		return nil, err
	}

	cfg = Config{
		DBConfig: DBConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
			DbSllMode:  os.Getenv("DB_SLL_MODE"),
		},
		HttpConfig: HttpConfig{
			ServerHost: os.Getenv("SERVER_HOST"),
			ServerPort: os.Getenv("SERVER_PORT"),
			AppName:    os.Getenv("APP_NAME"),
			AppHeader:  os.Getenv("APP_HEADER"),
		},
		FolderConfig: FolderConfig{
			RootPath:   os.Getenv("ROOT_PATH"),
			PublicPath: os.Getenv("PUBLIC_PATH"),
		},
	}

	return &cfg, nil
}
