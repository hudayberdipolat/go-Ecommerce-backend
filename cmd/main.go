package main

import (
	"fmt"
	"log"

	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/app"
)

func main() {
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal("error : ", err.Error())
	}
	// constructor

	// new app
	runSever := fmt.Sprintf("%s:%s", getDependencies.Config.HttpConfig.ServerHost, getDependencies.Config.HttpConfig.ServerPort)

	appRouter := app.NewApp(getDependencies)

	if errRunServer := appRouter.Listen(runSever); errRunServer != nil {
		log.Println("error : ", errRunServer.Error())
	}
}
