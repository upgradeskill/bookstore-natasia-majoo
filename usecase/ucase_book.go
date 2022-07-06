package usecase

import (
	"bookstore/dto"
	"bookstore/repository"
	"errors"

	"github.com/labstack/echo/v4"
)

type BookUsecase interface {
	FetchAll() ([]dto.Book, error)
	FetchByID(uint) (dto.Book, error)
	DeleteByID(uint) (dto.Book, error)
	Insert(ctx echo.Context) (dto.Book, error)
	Update(ctx echo.Context) (dto.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

var bookUsecaseInstance *bookUseCase

func InitBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	if bookUsecaseInstance == nil {
		bookUsecaseInstance = &bookUseCase{
			bookRepository: bookRepository,
		}
	}

	return bookUsecaseInstance
}

func (u *bookUseCase) FetchAll() (books []dto.Book, err error) {
	books, err = u.bookRepository.FetchAll()
	return
}

func (u *bookUseCase) FetchByID(ID uint) (book dto.Book, err error) {
	return u.bookRepository.FetchByID(ID)
}

func (u *bookUseCase) DeleteByID(ID uint) (book dto.Book, err error) {
	return u.bookRepository.DeleteByID(ID)
}

func (u *bookUseCase) Insert(ctx echo.Context) (book dto.Book, err error) {
	var payload dto.Book
	if err = ctx.Bind(&payload); err != nil {
		return
	}
	return u.bookRepository.Insert(payload)
}

func (u *bookUseCase) Update(ctx echo.Context) (book dto.Book, err error) {
	var payload dto.Book
	if err = ctx.Bind(&payload); err != nil {
		return
	}
	// check ID must be filled
	if payload.ID < 1 {
		return payload, errors.New("missing ID")
	}

	return u.bookRepository.Update(payload)
}
