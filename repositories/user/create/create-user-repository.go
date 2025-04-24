package repository

import (
	"api/initializers"
	"api/models"
)

func CreateUser(user *models.User) error {
	return initializers.DB.Create(user).Error
}
