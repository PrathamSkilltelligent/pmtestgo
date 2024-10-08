package config

import (
	"fmt"
	"syscall"
)

type AppConfig struct {
	// db           *databaseConfig
	server *serverConfig
	// jwt          *jwtConfig
	// loggerConfig *loggerConfig
	commonConfig *commonConfig
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {

	// dbConfig, err := newDBConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	serverConfig, err := newServerConfig()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}

	// jwtConfig, err := newJWTConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	// loggerConfig, err := newLoggerConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	commonConfig, err := newCommonConfig()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}

	appConfig = &AppConfig{
		// db:           dbConfig,
		server: serverConfig,
		// jwt:          jwtConfig,
		// loggerConfig: loggerConfig,
		commonConfig: commonConfig,
	}

	return appConfig
}
