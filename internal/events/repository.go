package events

import (
	"event-app/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]models.Event, error)
	GetByID(id string, loadCreatedBy bool, loadParticipants bool) (models.Event, error)
	Create(event models.Event) (models.Event, error)
	Update(event models.Event) (models.Event, error)
	Delete(id string) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := r.DB.Find(&events).Error
	return events, err
}

func (r *repository) Create(event models.Event) (models.Event, error) {
	err := r.DB.Create(&event).Error
	return event, err
}

func (r *repository) GetByID(id string, loadCreatedBy bool, loadParticipants bool) (models.Event, error) {
	var event models.Event
	query := r.DB
	if loadCreatedBy {
		query = query.Preload("User")
	}
	if loadParticipants {
		query = query.Preload("Participants.User")
	}

	err := query.First(&event, "id = ?", id).Error
	if err != nil {
		return models.Event{}, err
	}
	return event, err
}

func (r *repository) Update(event models.Event) (models.Event, error) {
	err := r.DB.Save(&event).Error
	return event, err
}

func (r *repository) Delete(id string) (bool, error) {
	err := r.DB.Delete(&models.Event{}, "id = ?", id).Error
	return err == nil, err
}
