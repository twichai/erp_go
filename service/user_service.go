package service

import (
	"erp/models"
	"erp/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}
