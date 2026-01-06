package middlewares

import (
	"event-app/internal/common"
	"event-app/pkg/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionMiddleware(db *gorm.DB, permissionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleIds, exists := c.Get("roleIds")
		if !exists {
			c.Error(common.ForbiddenError("User does not have roles"))
			c.Abort()
			return
		}

		hasPermission, err := security.HasPermission(db, roleIds.([]string), permissionName)
		if err != nil {
			c.Error(common.InternalServerError())
			c.Abort()
			return
		}

		if !hasPermission {
			c.Error(common.ForbiddenError("You do not have permission to perform this action"))
			c.Abort()
			return
		}

		c.Next()
	}
}
