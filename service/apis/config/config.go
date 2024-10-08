package config

import (
	"fmt"
	"syscall"
)

type AppConfig struct {
	server       *serverConfig
	commonConfig *commonConfig
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {

	serverConfig, err := newServerConfig()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}

	commonConfig, err := newCommonConfig()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}

	appConfig = &AppConfig{
		server:       serverConfig,
		commonConfig: commonConfig,
	}

	return appConfig
}
