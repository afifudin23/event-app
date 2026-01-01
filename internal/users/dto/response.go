package dto

import (
	"event-app/internal/models"
	"time"

	"github.com/google/uuid"
)

type EventInfo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	ID        uuid.UUID   `json:"id"`
	Fullname  string      `json:"fullname"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Events    []EventInfo `json:"events"`
}
type UserListItemResponse struct {
	ID        uuid.UUID `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Users []UserListItemResponse `json:"users"`
}

type SuccessResponse struct {
	ID uuid.UUID `json:"id"`
}

func ToResponse(user models.User) UserResponse {
	events := make([]EventInfo, 0, len(user.Events))
	for _, e := range user.Events {
		events = append(events, EventInfo{
			ID:        e.ID,
			Title:     e.Title,
			Location:  e.Location,
			StartDate: e.StartDate,
			EndDate:   e.EndDate,
			IsActive:  e.IsActive,
			CreatedAt: e.CreatedAt,
		})
	}

	return UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Events:    events,
	}
}

func ToListResponse(users []models.User) UserListResponse {
	var responses []UserListItemResponse
	for _, user := range users {
		responses = append(responses, UserListItemResponse{
			ID:        user.ID,
			Fullname:  user.Fullname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return UserListResponse{
		Users: responses,
	}
}

func ToSuccessResponse(id uuid.UUID) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}
