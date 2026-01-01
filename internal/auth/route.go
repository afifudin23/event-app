package auth

import (
	"event-app/internal/config"
	"event-app/internal/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := users.NewRepository(db)
	service := NewService(repo, cfg)
	handler := NewHandler(service, cfg)

	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}
}
