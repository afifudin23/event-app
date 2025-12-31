package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	users := r.Group("/users")
	{
		users.GET("/", handler.GetAllUsers)
		users.GET("/:id", handler.GetUserByID)
		users.POST("/", handler.CreateUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
