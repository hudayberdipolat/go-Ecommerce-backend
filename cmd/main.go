package main

import (
	"fmt"
	"log"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/setup/constructor"
)

func main() {
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal("error : ", err.Error())
	}
	constructor.Build(getDependencies)
	runSever := fmt.Sprintf("%s:%s", getDependencies.Config.HttpConfig.ServerHost, getDependencies.Config.HttpConfig.ServerPort)
	appRouter := app.NewApp(getDependencies)
	getDependencies.Logger.Info().Msg("Project run successfully")
	if errRunServer := appRouter.Listen(runSever); errRunServer != nil {
		getDependencies.Logger.Info().Msg("Project run server error")
		getDependencies.Logger.Info().Msg(errRunServer.Error())
	}
	// logger message

}
