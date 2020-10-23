package controller

import (
	c "breakfaster/config"
	"breakfaster/controller/v1/middleware"
	rv1 "breakfaster/controller/v1/router"
	"io"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server type for running the application
type Server struct {
	config  *c.Config
	engine  *gin.Engine
	auth    *middleware.AuthChecker
	routers map[string]interface{}
}

// NewEngine is a factory for gin engine instance
// Global Middlewares and api log configurations are registered here
func NewEngine(config *c.Config) *gin.Engine {
	gin.SetMode(config.GinMode)
	gin.DefaultWriter = io.MultiWriter(config.Logger.Writer)
	log.SetOutput(gin.DefaultWriter)

	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.PromMiddleware())
	return engine
}

// NewServer is the factory for server instance
func NewServer(config *c.Config, engine *gin.Engine, auth *middleware.AuthChecker, v1Router *rv1.Router) *Server {
	return &Server{
		config: config,
		engine: engine,
		auth:   auth,
		routers: map[string]interface{}{
			"v1": v1Router,
		},
	}
}

// RegisterRoutes method register all endpoints and returns a router
func (s *Server) RegisterRoutes() {
	s.engine.GET("/metrics", middleware.PromHandler(promhttp.Handler()))

	botGroup := s.engine.Group("/")
	{
		botGroup.POST("/callback", gin.WrapF(s.routers[s.config.BotVersion].(*(rv1.Router)).Bot.Callback))
	}

	v1 := s.engine.Group("/api/v1")
	{
		v1.GET("/foods", s.routers["v1"].(*(rv1.Router)).GetFoodAll)

		// TODO: JWT token for admin auth
		v1.GET("/employee/line-uid", s.routers["v1"].(*(rv1.Router)).GetEmployeeByEmpID)
		v1.GET("/employee/emp-id", s.routers["v1"].(*(rv1.Router)).GetEmployeeByLineUID)
		v1.POST("/employee", s.routers["v1"].(*(rv1.Router)).UpsertEmployeeByIDs)

		// TODO: JWT token for aunty auth
		v1.GET("/order", s.routers["v1"].(*(rv1.Router)).GetOrder)
		v1.PUT("/order/pick", s.routers["v1"].(*(rv1.Router)).SetPick)

		withAuth := v1.Group("/")
		withAuth.Use(s.auth.LineUIDAuth())
		{
			withAuth.POST("/orders", s.routers["v1"].(*(rv1.Router)).CreateOrders)
		}

		v1.GET("/next-week", s.routers["v1"].(*(rv1.Router)).GetNextWeekInterval)

		if mode := gin.Mode(); mode == gin.DebugMode {
			v1.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// Run is a method for starting server
func (s *Server) Run() {
	s.RegisterRoutes()

	Addr := ":" + s.config.Port
	s.engine.Run(Addr)
}

// BroadCastMenu is a method for broadcasting breakfast menu
func (s *Server) BroadCastMenu() error {
	if err := s.routers[s.config.BotVersion].(*(rv1.Router)).BroadCastMenu(); err != nil {
		return err
	}
	return nil
}
