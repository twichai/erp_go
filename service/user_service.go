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

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.UserRepo.GetUser(id)
}

func (s *UserService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	return s.UserRepo.UpdateUser(id, user)

}
