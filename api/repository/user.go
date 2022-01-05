package repository

import (
	"fuko_backend/infrastructure"
	"fuko_backend/models"
	"fuko_backend/utils"
)

// Accessing database
type UserRepository struct {
	db infrastructure.Database
}

// Instance on UserRepository
func FukoUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// Save user to database
func (u UserRepository) CreateUser(user models.UserRegister) error {
	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Phone = user.LastName
	dbUser.Sex = user.Sex
	dbUser.Username = user.Username
	dbUser.Password = user.Password
	dbUser.IsActive = true

	return u.db.DB.Create(&dbUser).Error

}

// LoginUser -> returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := utils.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
