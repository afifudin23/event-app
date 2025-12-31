package dto

import (
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
	CurrentID uuid.UUID      `json:"current_id"`
	Users     []UserResponse `json:"users"`
}

type SuccessResponse struct {
	ID uuid.UUID `json:"id"`
}
