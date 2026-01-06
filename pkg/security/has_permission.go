package security

import (
	"gorm.io/gorm"
)

func HasPermission(db *gorm.DB, roleIds []string, permissionName string) (bool, error) {
	var count int64

	err := db.Table("role_permissions").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id IN ? AND permissions.name = ?", roleIds, permissionName).
		Count(&count).Error

	return count > 0, err
}
