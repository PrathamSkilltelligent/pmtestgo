package config

// here AppContext is serverConfig (it inherits all variables and finctions of server config)
type AppContext struct {
	*serverConfig
}

func (ctx *AppContext) IsApplicationContext() {}

func (ctx *AppContext) WithAppConfig(appConfig *AppConfig) *AppContext {
	ctx.serverConfig = appConfig.server
	return ctx
}
