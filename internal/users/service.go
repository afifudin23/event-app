package users

import (
	"event-app/internal/common"
)

type Service interface {
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	Create(user User) (bool, error)
	Update(id string, user User) (bool, error)
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

func (s *service) Create(payload User) (bool, error) {
	return s.Repo.Create(payload)
}

func (s *service) Update(id string, payload User) (bool, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return false, common.NotFoundError("User not found")
	}

	user.Fullname = payload.Fullname
	user.Email = payload.Email
	user.Password = payload.Password
	return s.Repo.Update(*user)
}

func (s *service) Delete(id string) (bool, error) {
	if _, err := s.Repo.GetByID(id); err != nil {
		return false, common.NotFoundError("User not found")
	}
	return s.Repo.Delete(id)
}
