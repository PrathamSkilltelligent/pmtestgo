package config

// "bitbucket.org/_nkhalasi/lyra.go/logger"
// "github.com/nicksnyder/go-i18n/v2/i18n"

type AppContext struct {
	// *dbContext
	*serverConfig
	// *jwtConfig
	// *logger.Logger
	*commonConfig
	// messageBundle *i18n.Bundle
}

// func (ctx *AppContext) SetFaultBundle(mb *i18n.Bundle) {
// 	ctx.messageBundle = mb
// }

// func (ctx *AppContext) GetFaultBundle() *i18n.Bundle {
// 	return ctx.messageBundle
// }

// func (ctx *AppContext) SetLogger(logger *logger.Logger) {
// 	ctx.Logger = logger
// }

// func (ctx *AppContext) GetLogger() *logger.Logger {
// 	return ctx.Logger
// }

// func (ctx *AppContext) IsApplicationContext() {}

// func (ctx *AppContext) GetMessageBundle() *i18n.Bundle {
// 	return ctx.messageBundle
// }

func (ctx *AppContext) WithAppConfig(appConfig *AppConfig) *AppContext {
	// db, err := newDB(appConfig.db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	// msgBundle, err := getMessageBundle(db.GetDBObj())
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	// logger, err := newLogger(appConfig.loggerConfig.url, appConfig.server.IsProd())
	// if err != nil {
	// 	fmt.Println(err)
	// 	syscall.Exit(1)
	// }

	// ctx.dbContext = db
	// ctx.Logger = logger
	ctx.serverConfig = appConfig.server
	// ctx.jwtConfig = appConfig.jwt
	// ctx.messageBundle = msgBundle
	ctx.commonConfig = appConfig.commonConfig
	return ctx
}
