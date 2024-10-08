package config

type commonConfig struct {
	passPhrase string
}

func newCommonConfig() (*commonConfig, error) {
	passPhrase, err := GetEnvVar(PASSPHRASE)
	if err != nil {
		return nil, err
	}

	return &commonConfig{
		passPhrase: passPhrase,
	}, nil
}

func (c *commonConfig) GetPassPhrase() string {
	return c.passPhrase
}
