package auth

import (
	"event-app/internal/auth/dto"
	"event-app/internal/models"
	"event-app/internal/users"
)

func ToLoginResponse(user models.User, accessToken string) dto.UserLoginResponse {
	return dto.UserLoginResponse{
		User:        users.ToResponse(user),
		AccessToken: accessToken,
	}
}
func ToRegisterResponse(user models.User, accessToken string) dto.UserRegisterResponse {
	return dto.UserRegisterResponse{
		User:        users.ToResponse(user),
		AccessToken: accessToken,
	}
}
