package users

import (
	"event-app/internal/common"
	"event-app/internal/models"
	"event-app/internal/users/dto"
	"event-app/pkg/security"

	"github.com/google/uuid"
)

type Service interface {
	FindAll() ([]models.User, error)
	FindByID(id uuid.UUID) (models.User, error)
	Create(payload dto.UserRequest) (models.User, error)
	Update(id uuid.UUID, payload dto.UserRequest) (models.User, error)
	Delete(id uuid.UUID) (bool, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FindAll() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *service) FindByID(id uuid.UUID) (models.User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return models.User{}, common.NotFoundError("User not found")
	}
	return user, err
}

func (s *service) Create(payload dto.UserRequest) (models.User, error) {
	if _, err := s.Repo.GetByEmail(payload.Email); err == nil {
		return models.User{}, common.BadRequestError("Email already exists")
	}
	passwordHashed, err := security.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, err
	}
	return s.Repo.Create(models.User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: passwordHashed,
	})
}

func (s *service) Update(id uuid.UUID, payload dto.UserRequest) (models.User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return models.User{}, common.NotFoundError("User not found")
	}

	// CHECK DUPLICATE EMAIL
	existingEmail, err := s.Repo.GetByEmail(payload.Email)
	if err == nil && existingEmail.ID != user.ID {
		return models.User{}, common.BadRequestError("Email already exists")
	}

	// UPDATE
	passwordHashed, err := security.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Fullname = payload.Fullname
	user.Email = payload.Email
	user.Password = passwordHashed
	return s.Repo.Update(user)
}

func (s *service) Delete(id uuid.UUID) (bool, error) {
	if _, err := s.Repo.GetByID(id); err != nil {
		return false, common.NotFoundError("User not found")
	}
	return s.Repo.Delete(id)
}
