package middlewares

import (
	"event-app/internal/common"
	"log"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("PANIC:", r)

				common.ErrorHandler(c, common.InternalServerError())

				c.Abort()
			}
		}()

		c.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			common.ErrorHandler(c, c.Errors.Last().Err)
		}
	}
}
