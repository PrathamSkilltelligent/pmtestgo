package routers

import (
	// "boilerplate/common/log"
	"net/http"
	"os"

	lyraginerrors "github.com/PrathamSkilltelligent/pmgingo/errors"
	"github.com/PrathamSkilltelligent/pmgingo/routeutils"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/config"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/controllers/health"

	"github.com/gin-gonic/gin"
)

type Server struct {
	appContext *config.AppContext
	router     *gin.Engine
}

// here we are using gin.new() which creates basically the structure of out application and also the base route .
// but why not gin.Default() ?
// because gin.Default() actually creates some default logger and middleware that we dont need , we will build it out self .
func NewServer(appContext *config.AppContext) *Server {
	return &Server{appContext: appContext, router: gin.New()}
}

func (s *Server) AddRoutes() *Server {

	baseURL := s.router.Group("/api")
	baseURL.GET("/health", routeutils.HandleRequest(s.appContext, health.HealthController))
	return s
}

func (s *Server) Start() *http.Server {
	if s.appContext.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
	//here we are using http.server to create a server regarding gin provides us with a http.server using run() runction then why ?
	//here we are explicitely defining port from the appcontext and handler to gin.engine why?
	server := &http.Server{
		Addr:    ":" + s.appContext.GetPort(),
		Handler: s.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			f := lyraginerrors.ConfigError(err)
			_ = f
			os.Exit(1)
		}
	}()

	return server
}
