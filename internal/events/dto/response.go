package dto

import (
	"event-app/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserInfo struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
}

type EventResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   *UserInfo `json:"created_by,omitempty"`
}
type SuccessResponse struct {
	ID uuid.UUID `json:"id"`
}

func ToResponse(event models.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		Capacity:    event.Capacity,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		IsActive:    event.IsActive,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
		CreatedBy: &UserInfo{
			ID:       event.User.ID,
			Fullname: event.User.Fullname,
			Email:    event.User.Email,
		},
	}
}

func ToListResponse(events []models.Event) []EventResponse {
	var responses []EventResponse
	for _, event := range events {
		responses = append(responses, EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Location:    event.Location,
			Capacity:    event.Capacity,
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			IsActive:    event.IsActive,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
			CreatedBy:   nil,
		})
	}
	return responses
}

func ToSuccessResponse(id uuid.UUID) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}
