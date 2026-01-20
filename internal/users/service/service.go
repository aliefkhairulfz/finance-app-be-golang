package service

import (
	"finance-app-be/internal/users/model"
	"finance-app-be/internal/users/repository"
)

type Service struct {
	userRepo repository.RepositoryImpl
}

type ServiceImpl interface {
	FindAll() ([]model.User, error)
}

func NewService(userRepo repository.RepositoryImpl) ServiceImpl {
	return &Service{userRepo: userRepo}
}

func (s *Service) FindAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}
