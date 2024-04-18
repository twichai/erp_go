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

func (s *UserService) Login(user *models.User) (*models.User, error) {
	return s.UserRepo.Login(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.UserRepo.GetUser(id)
}
