package middlewares

import (
	"event-app/internal/common"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			common.ErrorHandler(c, c.Errors.Last().Err)
		}
	}
}
