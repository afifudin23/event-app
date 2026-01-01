package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Location    string    `gorm:"type:varchar(255);not null"`
	Capacity    int       `gorm:"type:integer;not null"`
	StartDate   time.Time `gorm:"type:timestampz;not null"`
	EndDate     time.Time `gorm:"type:timestampz;not null"`
	IsActive    bool      `gorm:"type:bool;not null"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null;references:users(id)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:CreatedBy;references:ID"`
}
