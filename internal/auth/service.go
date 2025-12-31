package auth

import (
	"event-app/internal/auth/dto"
	"event-app/internal/common"
	"event-app/internal/config"
	"event-app/internal/users"
)

type Service interface {
	Login(payload dto.UserLoginRequest) (*users.User, *string, error)
	Register(payload dto.UserRegisterRequest) (*users.User, *string, error)
}

type service struct {
	Repo users.Repository
	Cfg  *config.Config
}

func NewService(repo users.Repository, cfg *config.Config) Service {
	return &service{Repo: repo, Cfg: cfg}
}

func (s *service) Login(payload dto.UserLoginRequest) (*users.User, *string, error) {
	user, err := s.Repo.GetByEmail(payload.Email)
	if err != nil {
		return nil, nil, common.BadRequestError("User not found")
	}

	if !common.CheckPassword(payload.Password, user.Password) {
		return nil, nil, common.BadRequestError("Invalid password")
	}

	accessToken := common.GenerateToken(user.ID.String(), s.Cfg.SecretKey)

	return user, accessToken, nil
}
func (s *service) Register(payload dto.UserRegisterRequest) (*users.User, *string, error) {
	_, err := s.Repo.GetByEmail(payload.Email)
	if err == nil {
		return nil, nil, common.BadRequestError("Email already exists")
	}

	hashedPassword, err := common.HashPassword(payload.Password)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.Repo.Create(users.User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, nil, err
	}

	accessToken := common.GenerateToken(user.ID.String(), s.Cfg.SecretKey)

	return &user, accessToken, nil
}
