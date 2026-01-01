package events

import (
	"event-app/internal/common"
	"event-app/internal/events/dto"
	"event-app/internal/models"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	FindAll() ([]models.Event, error)
	FindByID(id uuid.UUID) (models.Event, error)
	Create(uid uuid.UUID, payload dto.EventRequest) (models.Event, error)
	Update(id uuid.UUID, payload dto.EventRequest) (models.Event, error)
	Delete(id uuid.UUID) (bool, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func validateEventDate(start, end time.Time) error {
	now := time.Now().UTC()

	if !start.After(now) {
		return common.BadRequestError("start_date must be greater than current time")
	}

	if !end.After(start) {
		return common.BadRequestError("end_date must be greater than start_date")
	}

	return nil
}

func (s *service) FindAll() ([]models.Event, error) {
	return s.Repo.GetAll()
}

func (s *service) FindByID(id uuid.UUID) (models.Event, error) {
	event, err := s.Repo.GetByID(id, true, true)
	if err != nil {
		return models.Event{}, common.NotFoundError("Event not found")
	}
	return event, err
}

func (s *service) Create(uid uuid.UUID, payload dto.EventRequest) (models.Event, error) {
	// if _, err := s.Repo.GetByTitle(payload.Title); err == nil {
	// 	return models.Event{}, common.BadRequestError("Title already exists")
	// }

	if err := validateEventDate(payload.StartDate, payload.EndDate); err != nil {
		return models.Event{}, err
	}

	return s.Repo.Create(models.Event{
		Title:       payload.Title,
		Description: payload.Description,
		Location:    payload.Location,
		Capacity:    payload.Capacity,
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		CreatedBy:   uid,
	})
}

func (s *service) Update(id uuid.UUID, payload dto.EventRequest) (models.Event, error) {
	event, err := s.Repo.GetByID(id, false, false)
	if err != nil {
		return models.Event{}, common.NotFoundError("Event not found")
	}

	// UPDATE
	event.Title = payload.Title
	event.Description = payload.Description
	event.Location = payload.Location
	event.Capacity = payload.Capacity
	event.StartDate = payload.StartDate
	event.EndDate = payload.EndDate

	return s.Repo.Update(event)
}

func (s *service) Delete(id uuid.UUID) (bool, error) {
	if _, err := s.Repo.GetByID(id, false, false); err != nil {
		return false, common.NotFoundError("Event not found")
	}
	return s.Repo.Delete(id)
}
