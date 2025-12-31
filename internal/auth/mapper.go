package auth

import (
	"event-app/internal/auth/dto"
	"event-app/internal/users"
)

func ToLoginResponse(user users.User, accessToken string) dto.UserLoginResponse {
	return dto.UserLoginResponse{
		User:        users.ToResponse(user),
		AccessToken: accessToken,
	}
}
