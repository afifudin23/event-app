package event_participants

import (
	"event-app/internal/config"
	"event-app/internal/middlewares"
	"event-app/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config, finder models.UserFinder) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	event_participant := r.Group("/events/:id")
	event_participant.Use(middlewares.AuthMiddleware(cfg.SecretKey, finder))
	{
		event_participant.GET("/participants", handler.GetAllEventParticipants)
		event_participant.POST("/register", handler.RegisterEventParticipant)
	}
}
