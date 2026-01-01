package dto

import (
	"event-app/internal/models"
	"event-app/internal/users/dto"
)

type UserLoginResponse struct {
	User        dto.UserResponse `json:"user"`
	AccessToken string           `json:"access_token"`
}

type UserRegisterResponse struct {
	User        dto.UserResponse `json:"user"`
	AccessToken string           `json:"access_token"`
}

func ToLoginResponse(user models.User, accessToken string) UserLoginResponse {
	return UserLoginResponse{
		User:        dto.ToResponse(user),
		AccessToken: accessToken,
	}
}
func ToRegisterResponse(user models.User, accessToken string) UserRegisterResponse {
	return UserRegisterResponse{
		User:        dto.ToResponse(user),
		AccessToken: accessToken,
	}
}
