package usecase

import (
	"bookstore/dto"
	"bookstore/repository"
	"bookstore/security"
	"errors"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Login(ctx echo.Context) (dto.UserLoginResponse, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
	jWtService     security.JWTService
}

var userUsecaseInstance *userUseCase

func InitUserUsecase(userRepository repository.UserRepository, jwtService security.JWTService) UserUsecase {
	if userUsecaseInstance == nil {
		userUsecaseInstance = &userUseCase{
			userRepository: userRepository,
			jWtService:     jwtService,
		}
	}

	return userUsecaseInstance
}

func (u *userUseCase) Login(ctx echo.Context) (response dto.UserLoginResponse, err error) {
	var payload dto.User
	if err = ctx.Bind(&payload); err != nil {
		return
	}
	if payload.Email == "" || payload.Password == "" {
		return response, errors.New("email or password must be filled")
	}
	response, err = u.userRepository.Login(payload.Email, payload.Password)
	if err != nil {
		return
	}
	response.Token, err = u.jWtService.GenerateToken(response.User.ID, response.User.Name)

	return
}
