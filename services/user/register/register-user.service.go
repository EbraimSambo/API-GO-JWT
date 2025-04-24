package services

import (
	"api/models"
	repository "api/repositories/user/create"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, email, password string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	err = repository.CreateUser(user)
	
	if err != nil {
		return nil, err
	}

	return user, nil
}
