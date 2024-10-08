package config

import (
	"fmt"
	"syscall"
)

const (
	JWT_ISSUER         = "JWT_ISSUER"
	JWT_SECRET         = "JWT_SECRET"
	JWT_PRIVATE_SECRET = "JWT_PRIVATE_SECRET"

	DB_HOSTNAME = "DB_HOSTNAME"
	DB_PORT     = "DB_PORT"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_TYPE     = "DB_TYPE"
	DB_NAME     = "DB_NAME"

	ENVIRONMENT = "ENVIRONMENT"
	APP_PORT    = "APP_PORT"

	LOG_HOST = "LOG_HOST"

	JWT_PRIVATE_KEY = "JWT_PRIVATE_KEY"
)

func GetEnvVar(varName string) (string, error) {
	val, found := syscall.Getenv(varName)
	if !found || val == "" {
		err := fmt.Errorf("env var %s not found", varName)
		return "", err
	} else {
		return val, nil
	}
}
