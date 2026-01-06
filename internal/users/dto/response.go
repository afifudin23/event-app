package dto

import (
	"event-app/internal/models"
	"time"
)

type EventInfo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserDetailResponse struct {
	ID             string      `json:"id"`
	Fullname       string      `json:"fullname"`
	Email          string      `json:"email"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Events         []EventInfo `json:"events"`
	Participations []EventInfo `json:"participations"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

type SuccessResponse struct {
	ID string `json:"id"`
}

func ToResponse(user models.User) UserResponse {
	roles := make([]string, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}

	return UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Roles:     roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToDetailResponse(user models.User) UserDetailResponse {
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
	participations := make([]EventInfo, 0, len(user.Participations))
	for _, e := range user.Participations {
		participations = append(participations, EventInfo{
			ID:        e.Event.ID,
			Title:     e.Event.Title,
			Location:  e.Event.Location,
			StartDate: e.Event.StartDate,
			EndDate:   e.Event.EndDate,
			IsActive:  e.Event.IsActive,
			CreatedAt: e.Event.CreatedAt,
		})
	}

	return UserDetailResponse{
		ID:             user.ID,
		Fullname:       user.Fullname,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Events:         events,
		Participations: participations,
	}
}

func ToListResponse(users []models.User) UserListResponse {
	var responses []UserResponse
	for _, user := range users {
		roles := make([]string, 0, len(user.Roles))
		for _, r := range user.Roles {
			roles = append(roles, r.Name)
		}
		responses = append(responses, UserResponse{
			ID:        user.ID,
			Fullname:  user.Fullname,
			Email:     user.Email,
			Roles:     roles,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return UserListResponse{
		Users: responses,
	}
}

func ToSuccessResponse(id string) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}
