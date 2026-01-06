package models

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;not null;references:roles(id)"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null;references:permissions(id)"`
}

type Role struct {
	ID         uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string       `gorm:"type:varchar(50);not null"`
	Permisions []Permission `gorm:"many2many:role_permissions"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
