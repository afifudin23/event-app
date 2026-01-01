package models

import (
	"time"

	"github.com/google/uuid"
)

type ParticipantStatus string

const (
	StatusRegistered ParticipantStatus = "registered"
	StatusCancelled  ParticipantStatus = "cancelled"
	StatusCheckedIn  ParticipantStatus = "checked_in"
)

type EventParticipants struct {
	ID      uuid.UUID         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	EventID uuid.UUID         `gorm:"type:uuid;not null;index;references:events(id)"`
	UserID  uuid.UUID         `gorm:"type:uuid;not null;index;references:users(id)"`
	Status  ParticipantStatus `gorm:"type:participant_status;default:'registered'"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Event Event `gorm:"foreignKey:EventID"`
	User  User  `gorm:"foreignKey:UserID"`
}
