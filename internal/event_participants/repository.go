package event_participants

import (
	"event-app/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(event_id string) ([]models.EventParticipants, error)
	Register(eventParticipants models.EventParticipants) (models.EventParticipants, error)
	Cancel(eventParticipants models.EventParticipants) (models.EventParticipants, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) FindAll(event_id string) ([]models.EventParticipants, error) {
	var eventParticipants []models.EventParticipants
	err := r.DB.
		Where("event_id = ?", event_id).
		// Preload("User").
		Find(&eventParticipants).
		Error
	return eventParticipants, err
}

func (r *repository) Register(eventParticipants models.EventParticipants) (models.EventParticipants, error) {
	err := r.DB.
		Where("event_id = ? AND user_id = ?", eventParticipants.EventID, eventParticipants.UserID).
		FirstOrCreate(&eventParticipants).
		Error
	return eventParticipants, err
}

func (r *repository) Cancel(eventParticipants models.EventParticipants) (models.EventParticipants, error) {
	if err := r.DB.Model(&models.EventParticipants{}).
		Where("event_id = ? AND user_id = ?", eventParticipants.EventID, eventParticipants.UserID).
		Update("status", eventParticipants.Status).Error; err != nil {
		return models.EventParticipants{}, err
	}

	var updated models.EventParticipants
	if err := r.DB.Where("event_id = ? AND user_id = ?", eventParticipants.EventID, eventParticipants.UserID).
		First(&updated).Error; err != nil {
		return models.EventParticipants{}, err
	}

	return updated, nil
}
