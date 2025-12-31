package root

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) {
	r.GET("", RootHandler)
}
