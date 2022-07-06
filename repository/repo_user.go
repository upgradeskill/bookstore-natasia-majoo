package repository

import (
	"bookstore/database"
	"bookstore/dto"
	"bookstore/util"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(string, string) (dto.UserLoginResponse, error)
}

type userRepo struct {
	DB *gorm.DB
}

func InitUserRepository() UserRepository {
	return &userRepo{
		DB: database.DB,
	}
}

func (repo *userRepo) Login(email string, password string) (response dto.UserLoginResponse, err error) {
	var user dto.User
	err = repo.DB.Find(&user, dto.User{Email: email}).Error
	if err != nil {
		return
	}
	if user.ID == 0 {
		return response, errors.New("user not found")
	}

	if err != nil || util.CheckPasswordHash(password, user.Password) {
		return response, errors.New("incorrect password")
	}
	response.User = user
	return
}
