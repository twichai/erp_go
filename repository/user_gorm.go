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
