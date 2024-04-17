package repository

import (
	"erp/models"

	"gorm.io/gorm"
)

type UserRepositoryGrom struct {
	DB gorm.DB
}

func (r *UserRepositoryGrom) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryGrom) Login(user *models.User) error {
	return r.DB.Where(&models.User{Username: user.Username, Password: user.Password}).First(&user).Error
}

func (r *UserRepositoryGrom) GetUser(id int) (*models.User, error) {
	user := models.User{}
	err := r.DB.First(&user, id).Error
	return &user, err
}
