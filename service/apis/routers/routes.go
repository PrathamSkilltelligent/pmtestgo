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

func NewServer(appContext *config.AppContext) *Server {
	return &Server{appContext: appContext, router: gin.New()}
}

func (s *Server) AddRoutes() *Server {

	baseURL := s.router.Group("/boilerplate/apis")
	baseURL.GET("/health", routeutils.HandleRequest(s.appContext, health.HealthController))
	return s
}

func (s *Server) Start() *http.Server {
	if s.appContext.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
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
