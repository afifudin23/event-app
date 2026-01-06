package models

import (
	"time"
)

type Event struct {
	ID          string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Location    string    `gorm:"type:varchar(255);not null"`
	Capacity    int       `gorm:"type:integer;not null"`
	StartDate   time.Time `gorm:"type:timestampz;not null"`
	EndDate     time.Time `gorm:"type:timestampz;not null"`
	IsActive    bool      `gorm:"type:bool;not null;default:true"`
	CreatedBy   string    `gorm:"type:uuid;not null;references:users(id)"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User         User                `gorm:"foreignKey:CreatedBy"`
	Participants []EventParticipants `gorm:"foreignKey:EventID"`
}
