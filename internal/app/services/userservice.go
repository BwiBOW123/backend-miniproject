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

func (s *UserService) Register(user *domain.Member) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByUsername(username string) (*domain.Member, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *UserService) Login(username, password string) (*domain.Member, error) {
	return s.repo.Login(username, password)
}
