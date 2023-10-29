package service

import (
	"BwiBOW123/backend-miniproject/internal/app/repository"
	"BwiBOW123/backend-miniproject/internal/domain"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByUsername(username string) (*domain.User, error) {
	return s.repo.GetUserByUsername(username)
}
