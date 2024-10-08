package config

import (
	"fmt"
	"syscall"
)

type AppConfig struct {
	server *serverConfig
}

func NewAppConfig() *AppConfig {

	// understood serverconfig
	serverConfig, err := newServerConfig()
	if err != nil {
		fmt.Println(err)
		//here we are using syscall.exit but in routes.go we are using os.exit why?
		syscall.Exit(1)
	}

	return &AppConfig{
		server: serverConfig,
	}

}
