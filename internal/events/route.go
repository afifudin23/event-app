package events

import (
	"event-app/internal/config"
	"event-app/internal/middlewares"
	"event-app/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config, finder models.UserFinder) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	events := router.Group("/events")
	events.Use(middlewares.AuthMiddleware(cfg.SecretKey, finder))

	{
		events.GET("", handler.GetAllEvents)
		events.POST("", handler.CreateEvent)
		events.GET("/:id", handler.GetEventByID)
		events.PUT("/:id", handler.UpdateEvent)
		events.DELETE("/:id", handler.DeleteEvent)
	}
}
