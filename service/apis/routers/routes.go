package routers

import (
	// "boilerplate/common/log"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/config"
	// "boilerplate/service/apis/controllers/auth"
	// "boilerplate/service/apis/controllers/health"
	// "boilerplate/service/apis/controllers/task"
	// "boilerplate/service/apis/middlewares"
	"net/http"
	"os"

	lyraginerrors "github.com/PrathamSkilltelligent/pmgingo/errors"
	// mw "bitbucket.org/_nkhalasi/lyra-gin.go/middlewares"

	"github.com/gin-gonic/gin"
)

// Server holds the necessary configurations for the server.
type Server struct {
	appContext *config.AppContext
	router     *gin.Engine
}

// NewServer creates a new instance of Server.
func NewServer(appContext *config.AppContext) *Server {
	return &Server{appContext: appContext, router: gin.New()}
}

// AddRoutes adds routes to the server.
func (s *Server) AddRoutes() *Server {

	baseURL := s.router.Group("/boilerplate/apis")
	_ = baseURL
	// baseURL.GET("/health", routeutils.HandleRequest(s.appContext, health.HealthController))

	// baseURL.POST("/signup", routeutils.HandleRequest(s.appContext, auth.SignupController))
	// baseURL.POST("/login", routeutils.HandleRequest(s.appContext, auth.LoginController))

	// v1 := baseURL.Group("/v1/orgs/:orgid")
	// v1.Use(routeutils.HandleMiddleware(s.appContext, middlewares.JWTVerificationMiddleware))
	// v1.POST("/task", routeutils.HandleRequest(s.appContext, task.CreateTaskController))
	// v1.GET("/tasks", routeutils.HandleRequest(s.appContext, task.GetTasksController))
	// v1.PATCH("/task", routeutils.HandleRequest(s.appContext, task.UpdateTaskController))
	// v1.PATCH("/task/complete", routeutils.HandleRequest(s.appContext, task.UpdateTaskCompleteController))
	// v1.PATCH("/task/:id/delete", routeutils.HandleRequest(s.appContext, task.DeleteTaskController))

	return s
}

// Start starts the HTTP server.
func (s *Server) Start() *http.Server {
	if s.appContext.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// s.router.Use(logmiddleware.GinMiddlewareLogger(s.appContext.Logger))
	// s.router.Use(mw.CORSMiddleware())
	// s.router.Use(mw.HSTSMiddleware())

	server := &http.Server{
		Addr:    ":" + s.appContext.GetPort(),
		Handler: s.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			f := lyraginerrors.ConfigError(err)
			_ = f
			// s.appContext.Logger.Error(log.NewLog(f.Error(), nil, nil, nil, nil, nil))
			os.Exit(1)
		}
	}()

	return server
}
