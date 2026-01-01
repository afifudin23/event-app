package event_participants

import (
	"event-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(event_id uuid.UUID) ([]models.EventParticipants, error)
	Register(eventParticipants models.EventParticipants) (models.EventParticipants, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) FindAll(event_id uuid.UUID) ([]models.EventParticipants, error) {
	var eventParticipants []models.EventParticipants
	err := r.DB.
		Where("event_id = ?", event_id).
		// Preload("User").
		Find(&eventParticipants).
		Error
	return eventParticipants, err
}

func (r *repository) Register(eventParticipants models.EventParticipants) (models.EventParticipants, error) {
	err := r.DB.Create(&eventParticipants).Error
	return eventParticipants, err
}
