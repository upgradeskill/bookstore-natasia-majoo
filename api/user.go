package api

import (
	"bookstore/repository"
	"bookstore/security"
	"bookstore/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorization(e *echo.Group) {
	var userRepo repository.UserRepository = repository.InitUserRepository()
	var jwtService security.JWTService = security.JWTAuthService()
	var userUseCase usecase.UserUsecase = usecase.InitUserUsecase(userRepo, jwtService)

	e.POST("/login", func(c echo.Context) error {
		data, err := userUseCase.Login(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	})
}
