package middlewares

import (
	"event-app/internal/common"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// CHECK AUTH HEADER
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(common.UnauthorizedError("Authorization header is required"))
			c.Abort()
			return
		}

		// GET BEARER TOKEN
		token := authHeader[len("Bearer "):]

		// VERIFY TOKEN
		claims, err := common.VerifyToken(token, secretKey)
		if err != nil {
			c.Error(common.UnauthorizedError(err.Error()))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
