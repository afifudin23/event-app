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
	user, err := s.Repo.GetByEmail(payload.Email, true)
	if err != nil {
		return models.User{}, "", common.BadRequestError("User not found")
	}

	if !security.CheckPassword(payload.Password, user.Password) {
		return models.User{}, "", common.BadRequestError("Invalid password")
	}

	roles, err := s.Repo.GetRolesByUserID(user.ID)
	if err != nil {
		return models.User{}, "", common.InternalServerError()
	}

	roleIDs := make([]string, 0)
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID.String())
	}

	accessToken := security.GenerateToken(user.ID, roleIDs, s.Cfg.SecretKey)

	return user, accessToken, nil
}

func (s *service) Register(payload dto.UserRegisterRequest) (models.User, string, error) {
	_, err := s.Repo.GetByEmail(payload.Email, false)
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

	// DEFAULT ROLE
	role, err := s.Repo.GetRoleByName("user")
	if err != nil {
		return models.User{}, "", common.InternalServerError()
	}

	if _, err = s.Repo.AssignRoleToUser(models.UserRole{
		UserID: user.ID,
		RoleID: role.ID.String(),
	}); err != nil {
		return models.User{}, "", common.InternalServerError()
	}

	roles, err := s.Repo.GetRolesByUserID(user.ID)
	if err != nil {
		return models.User{}, "", common.InternalServerError()
	}

	roleIDs := make([]string, 0)
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID.String())
	}
	user.Roles = roles

	accessToken := security.GenerateToken(user.ID, roleIDs, s.Cfg.SecretKey)

	return user, accessToken, nil
}
