package root

import "github.com/gin-gonic/gin"

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "ok",
	})
}
