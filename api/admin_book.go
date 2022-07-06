package api

import (
	"bookstore/repository"
	"bookstore/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AdminBookAPI(e *echo.Group) {
	var bookRepo repository.BookRepository = repository.InitBookRepository()
	var bookUseCase usecase.BookUsecase = usecase.InitBookUsecase(bookRepo)

	e.POST("/book", func(c echo.Context) error {
		data, err := bookUseCase.Insert(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

		}
		return c.JSON(http.StatusOK, data)
	})

	e.PUT("/book", func(c echo.Context) error {
		data, err := bookUseCase.Update(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

		}
		return c.JSON(http.StatusOK, data)
	})

	e.DELETE("/book/:ID", func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		data, err := bookUseCase.DeleteByID(uint(ID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

		}
		return c.JSON(http.StatusOK, data)
	})
}
