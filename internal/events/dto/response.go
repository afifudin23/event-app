package dto

import (
	"event-app/internal/models"
	"time"
)

type UserInfo struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type EventResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EventDetailResponse struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Location     string     `json:"location"`
	Capacity     int        `json:"capacity"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      time.Time  `json:"end_date"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Creator      UserInfo   `json:"creator"`
	Participants []UserInfo `json:"participants"`
}

type EventListResponse struct {
	Events []EventResponse `json:"events"`
}

type SuccessResponse struct {
	ID string `json:"id"`
}

func ToDetailResponse(event models.Event) EventDetailResponse {
	participants := make([]UserInfo, 0, len(event.Participants))

	for _, p := range event.Participants {
		participants = append(participants, UserInfo{
			ID:       p.User.ID,
			Fullname: p.User.Fullname,
			Email:    p.User.Email,
		})
	}

	return EventDetailResponse{
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
		Creator: UserInfo{
			ID:       event.User.ID,
			Fullname: event.User.Fullname,
			Email:    event.User.Email,
		},
		Participants: participants,
	}
}

func ToListResponse(events []models.Event) EventListResponse {
	responses := make([]EventResponse, 0, len(events))
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
		})
	}
	return EventListResponse{
		Events: responses,
	}
}

func ToSuccessResponse(id string) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}
