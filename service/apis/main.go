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

	var appContext = &config.AppContext{}
	appContext.WithAppConfig(appConfig)

	server := routers.NewServer(appContext).AddRoutes().Start()

	<-sutils.WaitForTermination(server)
}
