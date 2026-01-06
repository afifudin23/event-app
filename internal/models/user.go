package models

import (
	"time"
)

type User struct {
	ID       string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Fullname string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:text;not null"`
	Roles    []Role `gorm:"many2many:user_roles"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Events         []Event             `gorm:"foreignKey:CreatedBy"`
	Participations []EventParticipants `gorm:"foreignKey:UserID"`
}

type UserRole struct {
	UserID string `gorm:"type:uuid;not null;references:users(id)"`
	RoleID string `gorm:"type:uuid;not null;references:roles(id)"`
}
