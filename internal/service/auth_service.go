// internal/service/auth_service.go
package service

import (
	"errors"
	"go-auth/internal/model"
	"go-auth/internal/repository"
	"go-auth/pkg/jwt"
)

var userRepo = repository.NewUserRepository()

func RegisterUser(username, password string) error {
	user := model.User{UserName: username, Password: password}
	return userRepo.CreateUser(user)
}

func AuthenticateUser(username, password string) (string, error) {
	user, err := userRepo.GetUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.UserName)
	if err != nil {
		return "", err
	}
	return token, nil
}
