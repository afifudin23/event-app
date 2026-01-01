package event_participants

import (
	"event-app/internal/models"

	"github.com/google/uuid"
)

type Service interface {
	FindAll(event_id uuid.UUID) ([]models.EventParticipants, error)
	Register(uid uuid.UUID, event_id uuid.UUID) (models.EventParticipants, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FindAll(event_id uuid.UUID) ([]models.EventParticipants, error) {
	return s.Repo.FindAll(event_id)
}

func (s *service) Register(uid uuid.UUID, event_id uuid.UUID) (models.EventParticipants, error) {
	return s.Repo.Register(models.EventParticipants{
		EventID: event_id,
		UserID:  uid,
	})
}
