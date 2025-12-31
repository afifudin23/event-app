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
	Users []UserResponse `json:"users"`
}

type UserLoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}

type SuccessResponse struct {
	Created *bool `json:"created,omitempty"`
	Updated *bool `json:"updated,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
}