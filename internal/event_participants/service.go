package event_participants

import (
	"event-app/internal/models"
)

type Service interface {
	FindAll(event_id string) ([]models.EventParticipants, error)
	Register(uid string, event_id string) (models.EventParticipants, error)
	Cancel(uid string, event_id string) (models.EventParticipants, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FindAll(event_id string) ([]models.EventParticipants, error) {
	return s.Repo.FindAll(event_id)
}

func (s *service) Register(uid string, event_id string) (models.EventParticipants, error) {
	return s.Repo.Register(models.EventParticipants{
		EventID: event_id,
		UserID:  uid,
		Status:  models.ParticipantStatusRegistered,
	})
}

func (s *service) Cancel(uid string, event_id string) (models.EventParticipants, error) {
	return s.Repo.Cancel(models.EventParticipants{
		EventID: event_id,
		UserID:  uid,
		Status:  models.ParticipantStatusCancelled,
	})
}
