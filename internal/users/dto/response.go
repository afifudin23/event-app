package dto

import (
	"event-app/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

type SuccessResponse struct {
	ID uuid.UUID `json:"id"`
}

func ToResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListResponse(users []models.User) UserListResponse {
	var UserResponse []UserResponse
	for _, user := range users {
		UserResponse = append(UserResponse, ToResponse(user))
	}
	return UserListResponse{
		Users: UserResponse,
	}
}

func ToSuccessResponse(id uuid.UUID) SuccessResponse {
	return SuccessResponse{
		ID: id,
	}
}
