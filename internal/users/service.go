package users

import (
	"event-app/internal/common"
	"event-app/internal/users/dto"
)

type Service interface {
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	Create(payload dto.UserRequest) (User, error)
	Update(id string, payload dto.UserRequest) (User, error)
	Delete(id string) (bool, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FindAll() ([]User, error) {
	return s.Repo.GetAll()
}

func (s *service) FindByID(id string) (*User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, common.NotFoundError("User not found")
	}
	return user, err
}

func (s *service) Create(payload dto.UserRequest) (User, error) {
	if _, err := s.Repo.GetByEmail(payload.Email); err == nil {
		return User{}, common.BadRequestError("Email already exists")
	}
	hashedPassword, err := common.HashPassword(payload.Password)
	if err != nil {
		return User{}, err
	}
	return s.Repo.Create(User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: hashedPassword,
	})
}

func (s *service) Update(id string, payload dto.UserRequest) (User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return User{}, common.NotFoundError("User not found")
	}

	// CHECK DUPLICATE EMAIL
	existingEmail, err := s.Repo.GetByEmail(payload.Email)
	if err == nil && existingEmail.ID != user.ID {
		return User{}, common.BadRequestError("Email already exists")
	}

	// UPDATE
	hashedPassword, err := common.HashPassword(payload.Password)
	if err != nil {
		return User{}, err
	}
	user.Fullname = payload.Fullname
	user.Email = payload.Email
	user.Password = hashedPassword
	return s.Repo.Update(*user)
}

func (s *service) Delete(id string) (bool, error) {
	if _, err := s.Repo.GetByID(id); err != nil {
		return false, common.NotFoundError("User not found")
	}
	return s.Repo.Delete(id)
}
