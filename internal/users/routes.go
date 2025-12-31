package users

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

	users := r.Group("/users")
	users.Use(middlewares.AuthMiddleware(cfg.SecretKey))
	{
		users.GET("", handler.GetAllUsers)
		users.POST("", handler.CreateUser)
		users.GET("/:id", handler.GetUserByID)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
