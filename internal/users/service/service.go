package service

import (
	"finance-app-be/internal/users/model"
	"finance-app-be/internal/users/repository"
)

type Service struct {
	userRepo repository.RepositoryImpl
}

type ServiceImpl interface {
	FindAll() ([]model.UserResponse, error)
	FindOneById(id int) (*model.UserResponse, error)
	Create(email, password string) (*int64, error)
	UpdateOneById(id int, avatar *string, isActive *bool) (*int64, error)
}

func NewService(userRepo repository.RepositoryImpl) ServiceImpl {
	return &Service{userRepo: userRepo}
}

func (s *Service) FindOneById(id int) (*model.UserResponse, error) {
	return s.userRepo.FindOneById(id)
}

func (s *Service) FindAll() ([]model.UserResponse, error) {
	return s.userRepo.FindAll()
}

func (s *Service) Create(email, password string) (*int64, error) {
	return s.userRepo.Create(email, password)
}

func (s *Service) UpdateOneById(id int, avatar *string, isActive *bool) (*int64, error) {
	return s.userRepo.UpdateOneById(id, avatar, isActive)
}
