package repository

import (
	"api/initializers"
	"api/models"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}