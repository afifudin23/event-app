package users

import (
	"event-app/internal/models"
	"event-app/internal/users/dto"

	"github.com/google/uuid"
)

func ToResponse(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListResponse(users []models.User) dto.UserListResponse {
	var UserResponse []dto.UserResponse
	for _, user := range users {
		UserResponse = append(UserResponse, ToResponse(user))
	}
	return dto.UserListResponse{
		Users:     UserResponse,
	}
}

func ToSuccessResponse(id uuid.UUID) dto.SuccessResponse {
	return dto.SuccessResponse{
		ID: id,
	}
}
