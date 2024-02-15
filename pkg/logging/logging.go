package logging

import (
	"os"
	"strconv"

	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

func GetLogger(config *config.Config) *zerolog.Logger {
	folder := config.FolderConfig.RootPath + "logs"
	if err := os.MkdirAll(folder, 0777); err != nil {
		panic(err)
	}

	path := folder + "%Y-%m-%d.log"

	rl, _ := rotatelogs.New(path)

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	logger := zerolog.New(rl).With().Timestamp().Caller().Logger()
	return &logger
}
