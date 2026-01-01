package auth

import (
	"event-app/internal/auth/dto"
	"event-app/internal/common"
	"event-app/internal/config"
	"event-app/internal/models"
	"event-app/internal/users"
	"event-app/pkg/security"
)

type Service interface {
	Login(payload dto.UserLoginRequest) (models.User, string, error)
	Register(payload dto.UserRegisterRequest) (models.User, string, error)
}

type service struct {
	Repo users.Repository
	Cfg  *config.Config
}

func NewService(repo users.Repository, cfg *config.Config) Service {
	return &service{Repo: repo, Cfg: cfg}
}

func (s *service) Login(payload dto.UserLoginRequest) (models.User, string, error) {
	user, err := s.Repo.GetByEmail(payload.Email)
	if err != nil {
		return models.User{}, "", common.BadRequestError("User not found")
	}

	if !security.CheckPassword(payload.Password, user.Password) {
		return models.User{}, "", common.BadRequestError("Invalid password")
	}

	accessToken := security.GenerateToken(user.ID.String(), s.Cfg.SecretKey)

	return user, accessToken, nil
}
func (s *service) Register(payload dto.UserRegisterRequest) (models.User, string, error) {
	_, err := s.Repo.GetByEmail(payload.Email)
	if err == nil {
		return models.User{}, "", common.BadRequestError("Email already exists")
	}

	hashedPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, "", err
	}

	user, err := s.Repo.Create(models.User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return models.User{}, "", err
	}

	accessToken := security.GenerateToken(user.ID.String(), s.Cfg.SecretKey)

	return user, accessToken, nil
}
