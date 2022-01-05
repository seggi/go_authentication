package service

import (
	"fuko_backend/api/repository"
	"fuko_backend/models"
)

// UserService struct
type UserService struct {
	repo repository.UserRepository
}

// New user service -> get injected user repo
func FukoUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

// Save users entity
func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

// Login -> Gets validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
