package dto

import "event-app/internal/users/dto"

type UserLoginResponse struct {
	User        dto.UserResponse `json:"user"`
	AccessToken string           `json:"access_token"`
}

type UserRegisterResponse struct {
	User        dto.UserResponse `json:"user"`
	AccessToken string           `json:"access_token"`
}
