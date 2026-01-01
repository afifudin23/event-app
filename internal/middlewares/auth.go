package middlewares

import (
	"event-app/internal/common"
	"event-app/internal/models"
	"event-app/pkg/security"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware(secretKey string, finder models.UserFinder) gin.HandlerFunc {
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
		claims, err := security.VerifyToken(token, secretKey)
		if err != nil {
			c.Error(common.UnauthorizedError(err.Error()))
			c.Abort()
			return
		}

		// GET USER
		_, err = finder.FindByID(uuid.MustParse(claims.UID))
		if err != nil {
			c.Error(common.UnauthorizedError("User not found"))
			c.Abort()
			return
		}

		// SET USER ID
		c.Set("uid", claims.UID)
		c.Next()
	}
}
