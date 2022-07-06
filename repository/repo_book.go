package repository

import (
	"bookstore/database"
	"bookstore/dto"
	"time"

	"gorm.io/gorm"
)

type BookRepository interface {
	FetchAll() ([]dto.Book, error)
	FetchByID(uint) (dto.Book, error)
	DeleteByID(uint) (dto.Book, error)
	Insert(dto.Book) (dto.Book, error)
	Update(dto.Book) (dto.Book, error)
}

type bookRepo struct {
	DB *gorm.DB
}

// init repository. Called from api. Used to check if this repo already matched with the interface
func InitBookRepository() BookRepository {
	return &bookRepo{
		DB: database.DB,
	}
}

func (repo *bookRepo) FetchAll() (data []dto.Book, err error) {
	// get data
	err = repo.DB.Find(&data).Error
	return
}

func (repo *bookRepo) FetchByID(ID uint) (data dto.Book, err error) {
	// get book by id
	err = repo.DB.First(&data, ID).Error
	return
}

func (repo *bookRepo) DeleteByID(ID uint) (data dto.Book, err error) {
	// get book by id
	if err = repo.DB.First(&data, ID).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Delete(&data, ID).Error

	return
}

func (repo *bookRepo) Insert(payload dto.Book) (data dto.Book, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *bookRepo) Update(payload dto.Book) (data dto.Book, err error) {
	// get book by id
	if err = repo.DB.First(&data, payload.ID).Error; err != nil {
		return
	}

	// update value
	data.Name = payload.Name
	data.Author = payload.Author
	data.Description = payload.Description
	data.Publisher = payload.Publisher
	data.ISBN = payload.ISBN
	data.UpdatedAt = time.Now()

	// update book data
	err = repo.DB.Save(&data).Error
	return
}
