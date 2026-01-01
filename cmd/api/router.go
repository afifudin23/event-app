package api

import (
	"event-app/internal/auth"
	"event-app/internal/events"
	"event-app/internal/root"
	"event-app/internal/users"
)

func (s *Server) SetupRoutes() {
	v1 := s.Router.Group("/api/v1")
	userRepo := users.NewRepository(s.DB)
	userService := users.NewService(userRepo)
	{
		root.SetupRoutes(v1)
		auth.SetupRoutes(v1, s.DB, s.Cfg)
		users.SetupRoutes(v1, s.DB, s.Cfg)
		events.SetupRoutes(v1, s.DB, s.Cfg, userService)
	}
}
