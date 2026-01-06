package models

import (
	"time"
)

type ParticipantStatus string

const (
	ParticipantStatusRegistered ParticipantStatus = "registered"
	ParticipantStatusCancelled  ParticipantStatus = "cancelled"
	ParticipantStatusCheckedIn  ParticipantStatus = "checked_in"
)

type EventParticipants struct {
	ID      string            `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	EventID string            `gorm:"type:uuid;not null;index;references:events(id)"`
	UserID  string            `gorm:"type:uuid;not null;index;references:users(id)"`
	Status  ParticipantStatus `gorm:"type:participant_status;default:'registered'"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Event Event `gorm:"foreignKey:EventID"`
	User  User  `gorm:"foreignKey:UserID"`
}
