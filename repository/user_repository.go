package repository

import "erp/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
}
