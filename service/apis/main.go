package main

import (
	"fmt"

	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/config"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/routers"

	sutils "github.com/PrathamSkilltelligent/pmgingo/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env file not found")
	}

	appConfig := config.NewAppConfig()
	// fmt.Println(appConfig)
	var appContext = &config.AppContext{}
	appContext = appContext.WithAppConfig(appConfig)
	// fmt.Println(appContext)
	server := routers.NewServer(appContext).AddRoutes().Start()
	// server := routers.NewServer(appContext)
	// server1 := server.AddRoutes()
	// server2 := server1.Start()
	// fmt.Println(server)

	//i dont understand this below function and the flow click on it to expland
	<-sutils.WaitForTermination(server)
}
