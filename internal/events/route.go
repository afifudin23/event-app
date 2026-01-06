package events

import (
	"event-app/internal/config"
	"event-app/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	events := router.Group("/events")
	events.Use(middlewares.AuthMiddleware(db, cfg.SecretKey))

	{
		events.GET("", middlewares.PermissionMiddleware(db, "events.read"), handler.GetAllEvents)
		events.POST("", middlewares.PermissionMiddleware(db, "events.create"), handler.CreateEvent)
		events.GET("/:id", middlewares.PermissionMiddleware(db, "events.read"), handler.GetEventByID)
		events.PUT("/:id", middlewares.PermissionMiddleware(db, "events.update"), handler.UpdateEvent)
		events.DELETE("/:id", middlewares.PermissionMiddleware(db, "events.delete"), handler.DeleteEvent)
	}
}
