package api

import (
	"event-app/internal/root"
	"event-app/internal/users"
)

func (s *Server) SetupRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		root.SetupRoutes(v1)
		users.SetupRoutes(v1, s.DB)
	}
}
