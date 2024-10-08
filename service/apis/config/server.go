package config

type serverConfig struct {
	port         string
	isProduction bool
}

func newServerConfig() (*serverConfig, error) {

	port, err := GetEnvVar(APP_PORT)
	if err != nil {
		return nil, err
	}

	serverEnv, err := GetEnvVar(ENVIRONMENT)
	if err != nil {
		return nil, err
	}

	return &serverConfig{
		port:         port,
		isProduction: serverEnv == "prod",
	}, nil
}

func (config *serverConfig) GetPort() string {
	return config.port
}

func (config *serverConfig) IsProd() bool {
	return config.isProduction
}
