// internal/repository/user_repository.go
package repository

import (
	"go-auth/database"
	"go-auth/internal/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user model.User) error {
	return database.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
