package event_participants

import (
	"event-app/internal/config"
	"event-app/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	event_participant := r.Group("/events/:id")
	event_participant.Use(middlewares.AuthMiddleware(db, cfg.SecretKey))
	{
		event_participant.GET("/participants", middlewares.PermissionMiddleware(db, "event_participants.read"), handler.GetAllEventParticipants)
		event_participant.POST("/register", middlewares.PermissionMiddleware(db, "event_participants.register"), handler.RegisterEventParticipant)
		event_participant.POST("/cancel", middlewares.PermissionMiddleware(db, "event_participants.cancel"), handler.CancelEventParticipant)
	}
}
