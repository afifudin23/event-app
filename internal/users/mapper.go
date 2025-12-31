package users

import "event-app/internal/users/dto"

func ToResponse(user User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListResponse(users []User) dto.UserListResponse {
	var UserResponse []dto.UserResponse
	for _, user := range users {
		UserResponse = append(UserResponse, ToResponse(user))
	}
	return dto.UserListResponse{
		Users: UserResponse,
	}
}

func ToLoginResponse(user User, accessToken string) dto.UserLoginResponse {
	return dto.UserLoginResponse{
		User:        ToResponse(user),
		AccessToken: accessToken,
	}
}

func ToCreateResponse(success bool) dto.SuccessResponse {
	return dto.SuccessResponse{
		Created: &success,
	}
}

func ToUpdateResponse(success bool) dto.SuccessResponse {
	return dto.SuccessResponse{
		Updated: &success,
	}
}

func ToDeleteResponse(success bool) dto.SuccessResponse {
	return dto.SuccessResponse{
		Deleted: &success,
	}
}
