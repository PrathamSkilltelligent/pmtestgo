package config

type AppContext struct {
	*serverConfig
	*commonConfig
}

func (ctx *AppContext) IsApplicationContext() {}

func (ctx *AppContext) WithAppConfig(appConfig *AppConfig) *AppContext {
	ctx.serverConfig = appConfig.server
	ctx.commonConfig = appConfig.commonConfig
	return ctx
}
