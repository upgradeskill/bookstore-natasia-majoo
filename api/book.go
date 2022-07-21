package api

import (
	"bookstore/repository"
	"bookstore/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BookAPI(e *echo.Group) {
	var bookRepo repository.BookRepository = repository.InitBookRepository()
	var bookUseCase usecase.BookUsecase = usecase.InitBookUsecase(bookRepo)

	e.GET("/book", func(c echo.Context) error {
		data, err := bookUseCase.FetchAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/book/:ID", func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		data, err := bookUseCase.FetchByID(uint(ID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

		}
		return c.JSON(http.StatusOK, data)
	})

}
