package api

import (
	"event-app/internal/config"
	"event-app/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
	Cfg    *config.Config
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	router := gin.Default()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.Use(middlewares.ErrorMiddleware())
	server := &Server{
		Router: router,
		DB:     db,
		Cfg:    cfg,
	}

	server.SetupRoutes()

	return server
}

func (s *Server) Run() {
	s.Router.Run(":" + s.Cfg.Port)
}
