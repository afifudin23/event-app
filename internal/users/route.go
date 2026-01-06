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
	users.Use(middlewares.AuthMiddleware(db, cfg.SecretKey))
	{
		users.GET("", middlewares.PermissionMiddleware(db, "users.read"), handler.GetAllUsers)
		users.POST("", middlewares.PermissionMiddleware(db, "users.create"), handler.CreateUser)
		users.GET("/:id", middlewares.PermissionMiddleware(db, "users.read"), handler.GetUserByID)
		users.PUT("/:id", middlewares.PermissionMiddleware(db, "users.update"), handler.UpdateUser)
		users.DELETE("/:id", middlewares.PermissionMiddleware(db, "users.delete"), handler.DeleteUser)
	}
}
